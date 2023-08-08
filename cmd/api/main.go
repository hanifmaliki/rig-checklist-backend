package main

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/api/config"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/api/routes"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/error_handler"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

func main() {
	// Init log on debug mode as top priority
	logger.Init(0)

	// Load config
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg(err.Error())
	}
	appConfig := &config.Config{}
	err = envconfig.Process("", appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("on populate API config")
	}

	// Init log
	logger.Init(appConfig.Log.Level)

	app := fiber.New(fiber.Config{
		AppName:               "corporate-website-minerva api",
		DisableStartupMessage: true,
		BodyLimit:             5 * 1024 * 1024 * 1024,
		Concurrency:           256 * 1024,
		ErrorHandler:          error_handler.Handler,
	})

	// Init routes
	routes.Init(app, appConfig)

	log.Info().
		Msgf("Starting server on port %s", appConfig.Port)
	log.Fatal().
		Err(app.Listen(":"+appConfig.Port)).
		Msgf("Cannot start on port %s", appConfig.Port)
}
