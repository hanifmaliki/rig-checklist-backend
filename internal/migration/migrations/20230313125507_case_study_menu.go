package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125507 = gormigrate.Migration{
	ID: "20230313125507",
	Migrate: func(tx *gorm.DB) error {
		type CaseStudyMenu struct {
			model.GormAudit
			CaseStudyID int    `json:"case_study_id"`
			Type        string `json:"type"`
			Desc        string `json:"desc"`
			IsActive    bool   `json:"is_active"`
			OrderNo     int    `json:"order_no"`
		}
		return tx.AutoMigrate(&CaseStudyMenu{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("case_study_menus")
	},
}
