package migration

import (
	"os"

	"github.com/hanifmaliki/rig-checklist-backend/internal/migration/migrations"
	"github.com/hanifmaliki/rig-checklist-backend/internal/migration/seeds"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

func Init() {
	// Initialize Database
	db := persistence.Client().RigChecklist
	db.Logger.LogMode(logger.Info)

	// Starting migration of DB
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.Migrations)
	if err := m.Migrate(); err != nil {
		log.Fatal().
			Err(err).
			Msgf("could not migrate: %v", err)
	}
	log.Info().
		Msg("migration success")

	// Starting seeding of DB
	if os.Getenv("RUN_SEEDING") == "true" {
		for _, fn := range seeds.Seeds {
			if err := fn(db); err != nil {
				log.Fatal().
					Err(err).
					Msgf("could not seed: %v", err)
			}
		}
		log.Info().
			Msg("seeding success")
	}
}
