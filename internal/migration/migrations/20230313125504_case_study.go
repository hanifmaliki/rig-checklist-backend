package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125504 = gormigrate.Migration{
	ID: "20230313125504",
	Migrate: func(tx *gorm.DB) error {
		type CaseStudy struct {
			model.GormAudit
			Slug           string `json:"slug" gorm:"uniqueIndex:uidx_case_studies_slug;size:191"`
			CompanyName    string `json:"company_name"`
			Desc           string `json:"desc"`
			PersonName     string `json:"person_name"`
			PersonPosition string `json:"person_position"`
			Logo           string `json:"logo"`
			Banner         string `json:"banner"`
			ImpactDesc     string `json:"impact_desc"`
			IsActive       bool   `json:"is_active"`
			NotArchived    bool   `json:"not_archived" gorm:"->;type:boolean GENERATED ALWAYS AS (IF(deleted_at IS NULL, 1, NULL)) VIRTUAL;default:(-);uniqueIndex:uidx_case_studies_slug"`
		}
		return tx.AutoMigrate(&CaseStudy{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("case_studies")
	},
}
