package storage

import (
	"io"
)

type StorageInterface interface {
	Get(string, string) (io.Reader, error)
	GetPath(string, string) (string, error)
	Put(string, string, io.Reader) error
	Delete(string, string) error
}
