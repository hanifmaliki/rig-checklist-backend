package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124402_rig = gormigrate.Migration{
	ID: "20230811124402_rig",
	Migrate: func(tx *gorm.DB) error {
		type Rig struct {
			model.GormAudit
			Name     string `json:"name" gorm:"uniqueIndex:uidx_rigs_name,where:deleted_at IS NULL"`
			IsActive bool   `json:"is_active"`
		}
		return tx.AutoMigrate(&Rig{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("rigs")
	},
}
