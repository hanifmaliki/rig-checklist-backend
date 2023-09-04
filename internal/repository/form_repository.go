package repository

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"

	"gorm.io/gorm"
)

type FormRepository interface {
	CreateForm(data *model.Form) (*model.Form, error)
	ReadForms(conds map[string]interface{}, sortBy string) ([]*model.Form, error)
	ReadForm(id uint) (*model.Form, error)
	UpdateForm(data *model.Form) (*model.Form, error)
	DeleteForm(data *model.Form) error
}

func (r *repository) CreateForm(data *model.Form) (*model.Form, error) {
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

func (r *repository) ReadForm(id uint) (*model.Form, error) {
	db := persistence.Client().RigChecklist
	var data model.Form

	result := db.Preload("Rig").Preload("Location").Preload("Well").Preload("Field").Preload("Activity").Preload("Answers").First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadForms(conds map[string]interface{}, sortBy string) ([]*model.Form, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Form

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Preload("Rig").Preload("Location").Preload("Well").Preload("Field").Preload("Activity").Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateForm(formInput *model.Form) (*model.Form, error) {
	db := persistence.Client().RigChecklist
	var form model.Form

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&form, formInput.ID)
		if result.Error != nil {
			return result.Error
		}

		form.IsInspected1 = formInput.IsInspected1
		form.IsInspected2 = formInput.IsInspected2
		form.IsInspected3 = formInput.IsInspected3
		form.IsInspected4 = formInput.IsInspected4
		form.IsChecked1 = formInput.IsChecked1
		form.IsChecked2 = formInput.IsChecked2
		form.IsChecked3 = formInput.IsChecked3
		form.IsApproved = formInput.IsApproved
		form.Operable = formInput.Operable
		form.UpdatedBy = formInput.UpdatedBy

		result = tx.Save(&form)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &form, err
	}
}

func (r *repository) DeleteForm(formInput *model.Form) error {
	db := persistence.Client().RigChecklist
	var form model.Form

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&form, formInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&form).UpdateColumn("deleted_by", formInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&form)
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
