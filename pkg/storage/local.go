package storage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	directory string
	host      string
}

func InitLocal() (StorageInterface, error) {
	directory := fmt.Sprintf("%s/%s/upload", os.Getenv("GOPATH"), os.Getenv("BASE_PROJ_DIR"))
	host := os.Getenv("HOST")

	os.MkdirAll(directory, os.ModePerm)

	return Local{directory: directory, host: host}, nil
}

func (local Local) GetPath(filePrefix string, filename string) (string, error) {
	if filename == "" {
		return "", errors.New("file not found")
	}
	path := filepath.Join(local.host, "upload", filename)

	return path, nil
}

func (local Local) Get(filePrefix string, filename string) (io.Reader, error) {
	if filename == "" {
		return nil, errors.New("file not found")
	}
	path := filepath.Join(local.directory, filename)

	return os.Open(path)
}

func (local Local) Delete(filePrefix string, filename string) error {
	if filename == "" {
		return errors.New("file not found")
	}
	path := filepath.Join(local.directory, filename)

	return os.RemoveAll(path)
}

func (local Local) Put(filePrefix string, filename string, file io.Reader) error {
	path := filepath.Join(local.directory, filename)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)

	if fileInfo, err := os.Stat(path); fileInfo.Size() > 10000000 {
		if err != nil {
			return errors.New("file not found")
		}
		defer local.Delete(filePrefix, filename)
		return errors.New("file too large")
	}
	return nil
}
