package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type RigRepository interface {
	CreateRig(data *model.Rig) (*model.Rig, error)
	ReadRigs(conds map[string]interface{}, sortBy string) ([]*model.Rig, error)
	ReadRig(id uint) (*model.Rig, error)
	UpdateRig(data *model.Rig) (*model.Rig, error)
	DeleteRig(data *model.Rig) error
}

func (r *repository) CreateRig(data *model.Rig) (*model.Rig, error) {
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

func (r *repository) ReadRig(id uint) (*model.Rig, error) {
	db := persistence.Client().RigChecklist
	var data model.Rig

	result := db.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadRigs(conds map[string]interface{}, sortBy string) ([]*model.Rig, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Rig

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateRig(rigInput *model.Rig) (*model.Rig, error) {
	db := persistence.Client().RigChecklist
	var rig model.Rig

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&rig, rigInput.ID)
		if result.Error != nil {
			return result.Error
		}

		rig.Name = rigInput.Name
		rig.IsActive = rigInput.IsActive
		rig.UpdatedBy = rigInput.UpdatedBy

		result = tx.Save(&rig)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &rig, err
	}
}

func (r *repository) DeleteRig(rigInput *model.Rig) error {
	db := persistence.Client().RigChecklist
	var rig model.Rig

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&rig, rigInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&rig).UpdateColumn("deleted_by", rigInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&rig)
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
