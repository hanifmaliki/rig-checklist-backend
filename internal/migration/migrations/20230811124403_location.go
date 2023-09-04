package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124403_location = gormigrate.Migration{
	ID: "20230811124403_location",
	Migrate: func(tx *gorm.DB) error {
		type Location struct {
			model.GormAudit
			Name     string `json:"name" gorm:"uniqueIndex:uidx_locations_name,where:deleted_at IS NULL"`
			IsActive bool   `json:"is_active"`
		}
		return tx.AutoMigrate(&Location{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("locations")
	},
}
