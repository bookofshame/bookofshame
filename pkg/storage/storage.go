package storage

import "io"

type Storage interface {
	Upload(file io.Reader, fileName string) error
	UploadLarge(file io.Reader, fileName string) error
	Download(fileName string) ([]byte, error)
	DownloadLarge(fileName string) ([]byte, error)
	Delete(fileNames []string) error
}
