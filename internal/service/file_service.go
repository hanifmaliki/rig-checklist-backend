package service

import (
	"io"
	"mime/multipart"

	"github.com/rs/zerolog/log"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/minio"
)

type FileService interface {
	GetFileListFromMinio(prefix string) ([]*minio.FileDescription, error)
	GetFileFromMinio(filename string) (io.Reader, error)
	PutFileToMinio(filename string, file *multipart.FileHeader) error
	DeleteFileFromMinio(filename string) error
}

func (s *service) GetFileListFromMinio(prefix string) ([]*minio.FileDescription, error) {
	return minio.GetObjects(prefix)
}

func (s *service) GetFileFromMinio(filename string) (io.Reader, error) {
	return minio.GetObject(filename)
}

func (s *service) PutFileToMinio(filename string, file *multipart.FileHeader) error {
	object, err := file.Open()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	defer object.Close()

	err = minio.PutObject(filename, object, file.Size, file.Header.Get("Content-Type"))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (s *service) DeleteFileFromMinio(filename string) error {
	return minio.DeleteObject(filename)
}
