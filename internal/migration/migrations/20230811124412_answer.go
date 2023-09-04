package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124412_answer = gormigrate.Migration{
	ID: "20230811124412_answer",
	Migrate: func(tx *gorm.DB) error {
		type Answer struct {
			model.GormAudit
			FormID     uint   `json:"form_id" gorm:"uniqueIndex:uidx_answers,where:deleted_at IS NULL"`
			QuestionID uint   `json:"question_id" gorm:"uniqueIndex:uidx_answers,where:deleted_at IS NULL"`
			IsExist    *bool  `json:"is_exist"`
			IsGood     *bool  `json:"is_good"`
			NA         bool   `json:"na"`
			Note       string `json:"note"`
		}
		return tx.AutoMigrate(&Answer{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("answers")
	},
}
