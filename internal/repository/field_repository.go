package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type FieldRepository interface {
	CreateField(data *model.Field) (*model.Field, error)
	ReadFields(conds map[string]interface{}, sortBy string) ([]*model.Field, error)
	ReadField(id uint) (*model.Field, error)
	UpdateField(data *model.Field) (*model.Field, error)
	DeleteField(data *model.Field) error
}

func (r *repository) CreateField(data *model.Field) (*model.Field, error) {
	db := persistence.Client().RigChecklist
	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(data)
		if err := result.Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return data, err
	}
}

func (r *repository) ReadField(id uint) (*model.Field, error) {
	db := persistence.Client().RigChecklist
	var data model.Field

	result := db.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadFields(conds map[string]interface{}, sortBy string) ([]*model.Field, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Field

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateField(fieldInput *model.Field) (*model.Field, error) {
	db := persistence.Client().RigChecklist
	var field model.Field

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&field, fieldInput.ID)
		if result.Error != nil {
			return result.Error
		}

		field.Name = fieldInput.Name
		field.IsActive = fieldInput.IsActive
		field.UpdatedBy = fieldInput.UpdatedBy

		result = tx.Save(&field)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &field, err
	}
}

func (r *repository) DeleteField(fieldInput *model.Field) error {
	db := persistence.Client().RigChecklist
	var field model.Field

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&field, fieldInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&field).UpdateColumn("deleted_by", fieldInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&field)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}
