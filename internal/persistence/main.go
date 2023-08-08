package persistence

import (
	"regexp"
	"sync"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/db/mysql"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type clientList struct {
	Minerva *gorm.DB
	Petros  *gorm.DB
}

var (
	once    sync.Once
	clients = &clientList{}
)

func generateClient(config Config) *gorm.DB {
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatal().Err(err).Msg("on populate DB config")
	}

	var regex, _ = regexp.Compile(`\)\/|\?`)
	var dbName = regex.Split(config.GetDSN(), -1)[1]

	client := mysql.GetClient(config.GetDSN())
	log.Info().Msg("Connection to " + dbName + " DB opened")
	return client
}

func init() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Warn().Msg(err.Error())
		}
		clients.Minerva = generateClient(&MinervaConfig{})
		clients.Petros = generateClient(&PetrosConfig{})
	})
}

func Client() *clientList {
	return clients
}
