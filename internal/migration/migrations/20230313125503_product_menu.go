package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125503 = gormigrate.Migration{
	ID: "20230313125503",
	Migrate: func(tx *gorm.DB) error {
		type ProductMenu struct {
			model.GormAudit
			ProductID int    `json:"product_id"`
			Type      string `json:"type"`
			Desc      string `json:"desc"`
			IsActive  bool   `json:"is_active"`
			OrderNo   int    `json:"order_no"`
		}
		return tx.AutoMigrate(&ProductMenu{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("product_menus")
	},
}
