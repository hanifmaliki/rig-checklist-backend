package mysql

import (
	"github.com/rs/zerolog/log"
	gorm_zerolog "github.com/wei840222/gorm-zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetClient(dsn string) *gorm.DB {
	if dsn == "" {
		log.Fatal().
			Msg("empty mysql dsn")
	}

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_zerolog.New(),
	})

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("mysql.Connect")
	}
	return client
}
