package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125506 = gormigrate.Migration{
	ID: "20230313125506",
	Migrate: func(tx *gorm.DB) error {
		type CaseStudySolution struct {
			model.GormAudit
			CaseStudyID uint `json:"case_study_id"`
			ProductID   uint `json:"product_id"`
			OrderNo     int  `json:"order_no"`
		}
		return tx.AutoMigrate(&CaseStudySolution{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("case_study_solutions")
	},
}
