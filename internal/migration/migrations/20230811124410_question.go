package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124410_question = gormigrate.Migration{
	ID: "20230811124410_question",
	Migrate: func(tx *gorm.DB) error {
		type Question struct {
			model.GormAuditWithoutSoftDelete
			SectionID        uint   `json:"section_id" gorm:"uniqueIndex:uidx_questions"`
			SubSectionID     uint   `json:"sub_section_id" gorm:"uniqueIndex:uidx_questions"`
			No               string `json:"no" gorm:"uniqueIndex:uidx_questions"`
			SubNo            string `json:"sub_no" gorm:"uniqueIndex:uidx_questions"`
			Question         string `json:"question"`
			ClassificationID uint   `json:"classification_id"`
		}
		return tx.AutoMigrate(&Question{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("questions")
	},
}
