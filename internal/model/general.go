package model

import (
	"time"

	"gorm.io/gorm"
)

type GormAudit struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(254)"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(254)"`
	DeletedBy string         `json:"deleted_by" gorm:"type:varchar(254)"`
}

type GormAuditWithoutSoftDelete struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(254)"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(254)"`
}
