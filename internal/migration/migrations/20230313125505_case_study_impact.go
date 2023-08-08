package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125505 = gormigrate.Migration{
	ID: "20230313125505",
	Migrate: func(tx *gorm.DB) error {
		type CaseStudyImpact struct {
			model.GormAudit
			CaseStudyID int    `json:"case_study_id"`
			Title       string `json:"title"`
			SubTitle    string `json:"sub_title"`
			Desc        string `json:"desc"`
			Image       string `json:"image"`
			OrderNo     int    `json:"order_no"`
		}
		return tx.AutoMigrate(&CaseStudyImpact{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("case_study_impacts")
	},
}
