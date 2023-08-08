package main

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/migration"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	// Init log on debug mode as top priority
	logger.Init(0)

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg(err.Error())
	}

	// Migration engine
	migration.Init()
}
