package persistence

type Config interface {
	GetDSN() string
}

type RigChecklistConfig struct {
	DSN string `envconfig:"DB_KIMPER_POSTGRES_DSN" default:"host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"`
}

func (c RigChecklistConfig) GetDSN() string {
	return c.DSN
}
