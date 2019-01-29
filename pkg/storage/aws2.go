package storage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go"
)

type AWS2 struct {
	client *minio.Client
	opt    Option
}

type Option struct {
	Host     string
	Key      string
	Secret   string
	Secure   bool
	Bucket   string
	CdnHosts string
}

func InitAWS2() (StorageInterface, error) {
	awsOpt := Option{
		Host:     os.Getenv("RIAKCS_HOST"),
		Key:      os.Getenv("RIAKCS_KEY"),
		Secret:   os.Getenv("RIAKCS_SECRET"),
		Secure:   false,
		Bucket:   os.Getenv("RIAKCS_BUCKET"),
		CdnHosts: os.Getenv("CDN_HOSTS"),
	}
	client, err := minio.NewV2(awsOpt.Host, awsOpt.Key, awsOpt.Secret, awsOpt.Secure)
	if err != nil {
		log.Fatalln(err)
		return AWS2{}, err
	}

	exists, err := client.BucketExists(awsOpt.Bucket)
	if err != nil {
		log.Fatalln(err)
		return AWS2{}, err
	}

	if !exists {
		err = client.MakeBucket(awsOpt.Bucket, "us-east-1")
		if err != nil {
			log.Fatalln(err)
			return AWS2{}, err
		}
	}

	return AWS2{client: client, opt: awsOpt}, err
}

func (aws AWS2) Get(filePrefix string, filename string) (io.Reader, error) {
	if filename == "" {
		return nil, errors.New("file not found")
	}

	var objectName = filepath.Join(filePrefix, filename)
	return aws.client.GetObject(aws.opt.Bucket, objectName, minio.GetObjectOptions{})
}

func (aws AWS2) GetPath(filePrefix string, filename string) (string, error) {
	if filename == "" {
		return "", nil
	}

	hosts := strings.Split(aws.opt.CdnHosts, ",")
	host := hosts[random(0, len(hosts))]
	fileURL := fmt.Sprintf("%s/%s", host, filepath.Join(aws.opt.Bucket, filePrefix, filename))

	return fileURL, nil
}

func (aws AWS2) Delete(filePrefix string, filename string) error {
	if filename == "" {
		return errors.New("file not found")
	}

	err := aws.client.RemoveObject(aws.opt.Bucket, filepath.Join(filePrefix, filename))
	return err
}

func (aws AWS2) Put(filePrefix string, filename string, file io.Reader) error {
	ext := filepath.Ext(filename)
	ctype := mime.TypeByExtension(ext)
	_, err := aws.client.PutObject(aws.opt.Bucket, filepath.Join(filePrefix, filename), file, -1, minio.PutObjectOptions{ContentType: ctype, UserMetadata: map[string]string{"x-amz-acl": "public-read"}})
	return err
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
