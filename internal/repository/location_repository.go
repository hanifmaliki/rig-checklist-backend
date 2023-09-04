package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(data *model.Location) (*model.Location, error)
	ReadLocations(conds map[string]interface{}, sortBy string) ([]*model.Location, error)
	ReadLocation(id uint) (*model.Location, error)
	UpdateLocation(data *model.Location) (*model.Location, error)
	DeleteLocation(data *model.Location) error
}

func (r *repository) CreateLocation(data *model.Location) (*model.Location, error) {
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

func (r *repository) ReadLocation(id uint) (*model.Location, error) {
	db := persistence.Client().RigChecklist
	var data model.Location

	result := db.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadLocations(conds map[string]interface{}, sortBy string) ([]*model.Location, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Location

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateLocation(locationInput *model.Location) (*model.Location, error) {
	db := persistence.Client().RigChecklist
	var location model.Location

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&location, locationInput.ID)
		if result.Error != nil {
			return result.Error
		}

		location.Name = locationInput.Name
		location.IsActive = locationInput.IsActive
		location.UpdatedBy = locationInput.UpdatedBy

		result = tx.Save(&location)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &location, err
	}
}

func (r *repository) DeleteLocation(locationInput *model.Location) error {
	db := persistence.Client().RigChecklist
	var location model.Location

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&location, locationInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&location).UpdateColumn("deleted_by", locationInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&location)
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
