package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type WellRepository interface {
	CreateWell(data *model.Well) (*model.Well, error)
	ReadWells(conds map[string]interface{}, sortBy string) ([]*model.Well, error)
	ReadWell(id uint) (*model.Well, error)
	UpdateWell(data *model.Well) (*model.Well, error)
	DeleteWell(data *model.Well) error
}

func (r *repository) CreateWell(data *model.Well) (*model.Well, error) {
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

func (r *repository) ReadWell(id uint) (*model.Well, error) {
	db := persistence.Client().RigChecklist
	var data model.Well

	result := db.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadWells(conds map[string]interface{}, sortBy string) ([]*model.Well, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Well

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateWell(wellInput *model.Well) (*model.Well, error) {
	db := persistence.Client().RigChecklist
	var well model.Well

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&well, wellInput.ID)
		if result.Error != nil {
			return result.Error
		}

		well.Name = wellInput.Name
		well.IsActive = wellInput.IsActive
		well.UpdatedBy = wellInput.UpdatedBy

		result = tx.Save(&well)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &well, err
	}
}

func (r *repository) DeleteWell(wellInput *model.Well) error {
	db := persistence.Client().RigChecklist
	var well model.Well

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&well, wellInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&well).UpdateColumn("deleted_by", wellInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&well)
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
