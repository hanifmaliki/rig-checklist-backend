package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125502 = gormigrate.Migration{
	ID: "20230313125502",
	Migrate: func(tx *gorm.DB) error {
		type ProductHighlightedFeature struct {
			model.GormAudit
			ProductID int    `json:"product_id"`
			Name      string `json:"name"`
			Desc      string `json:"desc"`
			Image     string `json:"image"`
			OrderNo   int    `json:"order_no"`
		}
		return tx.AutoMigrate(&ProductHighlightedFeature{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("product_highlighted_features")
	},
}
