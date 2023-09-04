package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	CreateActivity(data *model.Activity) (*model.Activity, error)
	ReadActivities(conds map[string]interface{}, sortBy string) ([]*model.Activity, error)
	ReadActivity(id uint) (*model.Activity, error)
	UpdateActivity(data *model.Activity) (*model.Activity, error)
	DeleteActivity(data *model.Activity) error
}

func (r *repository) CreateActivity(data *model.Activity) (*model.Activity, error) {
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

func (r *repository) ReadActivity(id uint) (*model.Activity, error) {
	db := persistence.Client().RigChecklist
	var data model.Activity

	result := db.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadActivities(conds map[string]interface{}, sortBy string) ([]*model.Activity, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Activity

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateActivity(activityInput *model.Activity) (*model.Activity, error) {
	db := persistence.Client().RigChecklist
	var activity model.Activity

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&activity, activityInput.ID)
		if result.Error != nil {
			return result.Error
		}

		activity.Name = activityInput.Name
		activity.IsActive = activityInput.IsActive
		activity.UpdatedBy = activityInput.UpdatedBy

		result = tx.Save(&activity)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &activity, err
	}
}

func (r *repository) DeleteActivity(activityInput *model.Activity) error {
	db := persistence.Client().RigChecklist
	var activity model.Activity

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&activity, activityInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&activity).UpdateColumn("deleted_by", activityInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&activity)
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
