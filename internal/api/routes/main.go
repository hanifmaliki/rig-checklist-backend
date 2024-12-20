package routes

import (
	v1_router "github.com/hanifmaliki/rig-checklist-backend/internal/api/routes/v1"
	"github.com/hanifmaliki/rig-checklist-backend/pkg/healthz_router"
	"github.com/hanifmaliki/rig-checklist-backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hanifmaliki/rig-checklist-backend/internal/api/config"
	pk "github.com/hanifmaliki/rig-checklist-backend/pkg/private_key"
)

func Init(app fiber.Router, appConfig *config.Config) {
	api := app.Group("/minerva/api")
	healthz_router.HealthzRouter(api)
	api.Use(cors.New(cors.Config{
		ExposeHeaders: "Accept-Ranges, Content-Type, Content-Transfer-Encoding, Expires, Cache-Control, Pragma, File-Name, Content-Description, Content-Disposition",
	}))
	api.Use(requestid.New())
	api.Use(middleware.Logger())

	privateKey := pk.ReadPrivateKeyFile(appConfig.PrivateKeyPath)
	jwt := middleware.Jwt(privateKey, appConfig.DevMode)

	v1 := api.Group("/v1")
	v1_router.AuthRouter(v1.Group("/auth"), privateKey)
	v1_router.HiRouter(v1.Group("/hi"), &jwt)
	v1_router.FileRouter(v1.Group("/files"), &jwt)
}
