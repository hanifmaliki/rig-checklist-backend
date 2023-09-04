package migrations

import (
	"time"

	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230811124411_form = gormigrate.Migration{
	ID: "20230811124411_form",
	Migrate: func(tx *gorm.DB) error {
		type Form struct {
			model.GormAudit
			FormNo       string    `json:"form_no" gorm:"uniqueIndex:uidx_forms_form_no,where:deleted_at IS NULL"`
			RigID        uint      `json:"rig_id" gorm:"uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			LocationID   uint      `json:"location_id" gorm:"type:date;uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			WellID       uint      `json:"well_id" gorm:"type:date;uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			FieldID      uint      `json:"field_id" gorm:"type:date;uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			ActivityID   uint      `json:"activity_id" gorm:"type:date;uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			Date         time.Time `json:"date" gorm:"type:date;uniqueIndex:uidx_forms_identity,where:deleted_at IS NULL"`
			IsInspected1 bool      `json:"is_inspected_1"`
			IsInspected2 bool      `json:"is_inspected_2"`
			IsInspected3 bool      `json:"is_inspected_3"`
			IsInspected4 bool      `json:"is_inspected_4"`
			IsChecked1   bool      `json:"is_checked_1"`
			IsChecked2   bool      `json:"is_checked_2"`
			IsChecked3   bool      `json:"is_checked_3"`
			IsApproved   bool      `json:"is_approved"`
			Operable     *bool     `json:"operable"`
		}
		return tx.AutoMigrate(&Form{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("forms")
	},
}
