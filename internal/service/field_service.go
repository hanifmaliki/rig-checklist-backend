package service

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
)

type FieldService interface {
	CreateField(user *model.User, data *model.Field) (*model.Field, error)
	ReadFields(conds map[string]interface{}, sortBy string) ([]*model.Field, error)
	ReadField(id uint) (*model.Field, error)
	UpdateField(user *model.User, id uint, data *model.Field) (*model.Field, error)
	DeleteField(user *model.User, id uint) error
}

func (s *service) CreateField(user *model.User, data *model.Field) (*model.Field, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateField(data)
}

func (s *service) ReadFields(conds map[string]interface{}, sortBy string) ([]*model.Field, error) {
	return s.repository.ReadFields(conds, sortBy)
}

func (s *service) ReadField(id uint) (*model.Field, error) {
	return s.repository.ReadField(id)
}

func (s *service) UpdateField(user *model.User, id uint, data *model.Field) (*model.Field, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateField(data)
}

func (s *service) DeleteField(user *model.User, id uint) error {
	var data model.Field

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteField(&data)
}
