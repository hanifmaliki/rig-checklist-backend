package persistence

type Config interface {
	GetDSN() string
}

// type MinervaConfig struct {
// 	DSN string `envconfig:"DB_PETROS_MINERVA_POSTGRES_DSN" default:"host=127.0.0.1 user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"`
// }

type MinervaConfig struct {
	DSN string `envconfig:"DB_PETROS_MINERVA_MYSQL_DSN" default:"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"`
}

func (c MinervaConfig) GetDSN() string {
	return c.DSN
}

type PetrosConfig struct {
	DSN string `envconfig:"DB_PETROS_MYSQL_DSN" default:"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"`
}

func (c PetrosConfig) GetDSN() string {
	return c.DSN
}
