package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125500 = gormigrate.Migration{
	ID: "20230313125500",
	Migrate: func(tx *gorm.DB) error {
		type HomeContent struct {
			model.GormAuditWithoutSoftDelete
			Section string `json:"section" gorm:"uniqueIndex:uidx_home_contents;size:191"`
			Key     string `json:"key" gorm:"uniqueIndex:uidx_home_contents;size:191"`
			Value   string `json:"value"`
			IsJson  bool   `json:"is_json"`
		}
		return tx.AutoMigrate(&HomeContent{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("home_contents")
	},
}
