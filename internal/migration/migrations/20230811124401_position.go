package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124401_position = gormigrate.Migration{
	ID: "20230811124401_position",
	Migrate: func(tx *gorm.DB) error {
		type Position struct {
			model.GormAuditWithoutSoftDelete
			Name string `json:"name" gorm:"uniqueIndex:uidx_positions_name,where:deleted_at IS NULL"`
			Role string `json:"role"`
		}
		return tx.AutoMigrate(&Position{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("positions")
	},
}
