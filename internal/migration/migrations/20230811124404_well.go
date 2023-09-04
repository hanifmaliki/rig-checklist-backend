package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124404_well = gormigrate.Migration{
	ID: "20230811124404_well",
	Migrate: func(tx *gorm.DB) error {
		type Well struct {
			model.GormAudit
			Name     string `json:"name" gorm:"uniqueIndex:uidx_wells_name,where:deleted_at IS NULL"`
			IsActive bool   `json:"is_active"`
		}
		return tx.AutoMigrate(&Well{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("wells")
	},
}
