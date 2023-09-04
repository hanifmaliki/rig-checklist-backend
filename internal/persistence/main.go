package persistence

import (
	"sync"

	"github.com/hanifmaliki/rig-checklist-backend/pkg/db/postgres"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type clientList struct {
	RigChecklist *gorm.DB
}

var (
	once    sync.Once
	clients = &clientList{}
)

func generateClientPostgres(config Config) *gorm.DB {
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatal().Err(err).Msg("on populate DB config")
	}

	client := postgres.GetClient(config.GetDSN())
	log.Info().Msg("Connection to DB opened")
	return client
}

func init() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Warn().Msg(err.Error())
		}
		clients.RigChecklist = generateClientPostgres(&RigChecklistConfig{})
	})
}

func Client() *clientList {
	return clients
}
