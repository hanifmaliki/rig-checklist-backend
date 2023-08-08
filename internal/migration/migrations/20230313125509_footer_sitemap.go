package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125509 = gormigrate.Migration{
	ID: "20230313125509",
	Migrate: func(tx *gorm.DB) error {
		type FooterSitemap struct {
			model.GormAudit
			Name    string `json:"name"`
			Url     string `json:"url"`
			OrderNo int    `json:"order_no"`
		}
		return tx.AutoMigrate(&FooterSitemap{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("footer_sitemaps")
	},
}
