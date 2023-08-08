package minio

type Config struct {
	Endpoint        string `envconfig:"MINIO_ENDPOINT" default:"play.min.io"`
	AccessKeyID     string `envconfig:"MINIO_ACCESS_KEY_ID" default:"Q3AM3UQ867SPQQA43P2F"`
	SecretAccessKey string `envconfig:"MINIO_SECRET_ACCESS_KEY" default:"zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"`
	UseSSL          string `envconfig:"MINIO_USE_SSL"`
	BucketName      string `envconfig:"MINIO_BUCKET_NAME"`
}

type FileDescription struct {
	ContentType string `json:"content_type"`
	Name        string `json:"file_name"`
}
