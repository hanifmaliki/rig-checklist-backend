package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124409_classification = gormigrate.Migration{
	ID: "20230811124409_classification",
	Migrate: func(tx *gorm.DB) error {
		type Classification struct {
			model.GormAuditWithoutSoftDelete
			Name string `json:"name" gorm:"uniqueIndex:uidx_classifications_name"`
		}

		return tx.AutoMigrate(&Classification{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("classifications")
	},
}
