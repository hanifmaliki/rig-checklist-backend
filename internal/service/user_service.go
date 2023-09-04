package service

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/helper"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
)

type UserService interface {
	CreateUser(user *model.User, data *model.User) (*model.User, error)
	ReadUsers(conds map[string]interface{}, sortBy string) ([]*model.User, error)
	ReadUser(conds map[string]interface{}) (*model.User, error)
	UpdateUser(user *model.User, id uint, data *model.User) (*model.User, error)
	DeleteUser(user *model.User, id uint) error
}

func (s *service) CreateUser(user *model.User, data *model.User) (*model.User, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = hashedPassword

	return s.repository.CreateUser(data)
}

func (s *service) ReadUsers(conds map[string]interface{}, sortBy string) ([]*model.User, error) {
	return s.repository.ReadUsers(conds, sortBy)
}

func (s *service) ReadUser(conds map[string]interface{}) (*model.User, error) {
	return s.repository.ReadUser(conds)
}

func (s *service) UpdateUser(user *model.User, id uint, data *model.User) (*model.User, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateUser(data)
}

func (s *service) DeleteUser(user *model.User, id uint) error {
	var data model.User

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteUser(&data)
}
