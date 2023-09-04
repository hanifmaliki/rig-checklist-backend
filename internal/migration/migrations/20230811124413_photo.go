package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124413_photo = gormigrate.Migration{
	ID: "20230811124413_photo",
	Migrate: func(tx *gorm.DB) error {
		type Photo struct {
			model.GormAudit
			AnswerID uint   `json:"answer_id"`
			Url      string `json:"url"`
		}
		return tx.AutoMigrate(&Photo{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("photos")
	},
}
