package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125508 = gormigrate.Migration{
	ID: "20230313125508",
	Migrate: func(tx *gorm.DB) error {
		type Footer struct {
			model.GormAuditWithoutSoftDelete
			Address string `json:"address"`
			Phone   string `json:"phone"`
			Email   string `json:"email"`
		}
		return tx.AutoMigrate(&Footer{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("footers")
	},
}
