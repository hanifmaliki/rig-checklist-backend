package migrations

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124400_user = gormigrate.Migration{
	ID: "20230811124400_user",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.GormAudit
			Name       string `json:"name"`
			Username   string `json:"username" gorm:"uniqueIndex:uidx_users_username,where:deleted_at IS NULL"`
			Email      string `json:"email" gorm:"uniqueIndex:uidx_users_email,where:deleted_at IS NULL"`
			Password   string `json:"password"`
			PositionID uint   `json:"position_id"`
			FieldID    uint   `json:"field_id"`
			IsAdmin    bool   `json:"is_admin"`
			IsActive   bool   `json:"is_active"`
		}
		return tx.AutoMigrate(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
