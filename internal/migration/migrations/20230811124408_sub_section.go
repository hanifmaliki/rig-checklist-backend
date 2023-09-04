package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124408_sub_section = gormigrate.Migration{
	ID: "20230811124408_sub_section",
	Migrate: func(tx *gorm.DB) error {
		type SubSection struct {
			model.GormAuditWithoutSoftDelete
			SectionID uint   `json:"section_id" gorm:"uniqueIndex:uidx_sub_sections"`
			Name      string `json:"name" gorm:"uniqueIndex:uidx_sub_sections"`
		}
		return tx.AutoMigrate(&SubSection{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("sub_sections")
	},
}
