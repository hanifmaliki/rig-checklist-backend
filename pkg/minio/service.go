package minio

import (
	"context"
	"io"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

var (
	once        sync.Once
	minioClient *minio.Client
	minioConfig *Config
)

func init() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Warn().Msg(err.Error())
		}
		minioConfigTemp := &Config{}
		err = envconfig.Process("", minioConfigTemp)
		if err != nil {
			log.Fatal().Err(err).Msg("on populate Minio config")
		}
		minioConfig = minioConfigTemp
		var ssl bool
		if minioConfig.UseSSL != "" {
			ssl = true
		}
		useSSL := ssl

		// Initialize Minio client object
		minioClient, err = minio.New(minioConfig.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.SecretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("on initialize Minio client object")
		} else {
			log.Info().Msg("Initialize Minio client object success")
		}
	})
}

func GetObjects(prefix string) ([]*FileDescription, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objects := minioClient.ListObjects(ctx, minioConfig.BucketName, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	var results []*FileDescription
	for object := range objects {
		if object.Err != nil {
			log.Error().Msg(object.Err.Error())
			return nil, object.Err
		}

		objectInfo, _ := minioClient.StatObject(ctx, minioConfig.BucketName, object.Key, minio.GetObjectOptions{})
		results = append(results, &FileDescription{
			ContentType: objectInfo.ContentType,
			Name:        object.Key,
		})
	}

	return results, nil
}

func GetObject(filename string) (*minio.Object, error) {
	object, err := minioClient.GetObject(context.Background(), minioConfig.BucketName, filename, minio.GetObjectOptions{
		ServerSideEncryption: nil,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return object, nil
}

func PutObject(filename string, file io.Reader, fileSize int64, contentType string) error {
	info, err := minioClient.PutObject(context.Background(), minioConfig.BucketName, filename, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	log.Info().Msgf("File with title %s of size %d has been uploaded", info.ETag, info.Size)
	return nil
}

func FPutObject(filename string, filePath string) error {
	info, err := minioClient.FPutObject(context.Background(), minioConfig.BucketName, filename, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	log.Info().Msgf("File with title %s of size %d has been uploaded", info.ETag, info.Size)
	return nil
}

func DeleteObject(filename string) error {
	return minioClient.RemoveObject(context.Background(), minioConfig.BucketName, filename, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
}
