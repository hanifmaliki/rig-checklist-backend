package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data *model.User) (*model.User, error)
	ReadUsers(conds map[string]interface{}, sortBy string) ([]*model.User, error)
	ReadUser(conds map[string]interface{}) (*model.User, error)
	UpdateUser(data *model.User) (*model.User, error)
	DeleteUser(data *model.User) error
}

func (r *repository) CreateUser(data *model.User) (*model.User, error) {
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

func (r *repository) ReadUser(conds map[string]interface{}) (*model.User, error) {
	db := persistence.Client().RigChecklist
	var data model.User

	result := db.Preload("Role")
	if conds["id"] != nil {
		result = result.Where("id = ?", conds["id"])
	}
	if conds["username"] != nil {
		result = result.Where(`lower(username) = ? OR lower(email) = ?`,
			conds["username"], conds["username"])
	}
	result = result.First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadUsers(conds map[string]interface{}, sortBy string) ([]*model.User, error) {
	db := persistence.Client().RigChecklist
	var data []*model.User

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Preload("Role").Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateUser(userInput *model.User) (*model.User, error) {
	db := persistence.Client().RigChecklist
	var user model.User

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&user, userInput.ID)
		if result.Error != nil {
			return result.Error
		}

		user.Name = userInput.Name
		user.Username = userInput.Username
		user.Email = userInput.Email
		user.Password = userInput.Password
		user.PositionID = userInput.PositionID
		user.FieldID = userInput.FieldID
		user.IsAdmin = userInput.IsAdmin
		user.IsActive = userInput.IsActive
		user.UpdatedBy = userInput.UpdatedBy

		result = tx.Save(&user)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &user, err
	}
}

func (r *repository) DeleteUser(userInput *model.User) error {
	db := persistence.Client().RigChecklist
	var user model.User

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&user, userInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&user).UpdateColumn("deleted_by", userInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&user)
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
