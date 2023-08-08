package postgres

import (
	"github.com/rs/zerolog/log"
	gorm_zerolog "github.com/wei840222/gorm-zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClient(dsn string) *gorm.DB {
	if dsn == "" {
		log.Fatal().
			Msg("empty postgres dsn")
	}

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_zerolog.New(),
	})

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("postgres.Connect")
	}
	return client
}
