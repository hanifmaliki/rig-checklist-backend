package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124414_sign = gormigrate.Migration{
	ID: "20230811124414_sign",
	Migrate: func(tx *gorm.DB) error {
		type Sign struct {
			model.GormAudit
			UserID uint   `json:"user_id"`
			Url    string `json:"url"`
		}
		return tx.AutoMigrate(&Sign{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("signs")
	},
}
