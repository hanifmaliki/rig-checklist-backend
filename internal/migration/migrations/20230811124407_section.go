package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124407_section = gormigrate.Migration{
	ID: "20230811124407_section",
	Migrate: func(tx *gorm.DB) error {
		type Section struct {
			model.GormAuditWithoutSoftDelete
			Name string `json:"name" gorm:"uniqueIndex:uidx_sections_name"`
		}
		return tx.AutoMigrate(&Section{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("sections")
	},
}
