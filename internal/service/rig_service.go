package service

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
)

type RigService interface {
	CreateRig(user *model.User, data *model.Rig) (*model.Rig, error)
	ReadRigs(conds map[string]interface{}, sortBy string) ([]*model.Rig, error)
	ReadRig(id uint) (*model.Rig, error)
	UpdateRig(user *model.User, id uint, data *model.Rig) (*model.Rig, error)
	DeleteRig(user *model.User, id uint) error
}

func (s *service) CreateRig(user *model.User, data *model.Rig) (*model.Rig, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateRig(data)
}

func (s *service) ReadRigs(conds map[string]interface{}, sortBy string) ([]*model.Rig, error) {
	return s.repository.ReadRigs(conds, sortBy)
}

func (s *service) ReadRig(id uint) (*model.Rig, error) {
	return s.repository.ReadRig(id)
}

func (s *service) UpdateRig(user *model.User, id uint, data *model.Rig) (*model.Rig, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateRig(data)
}

func (s *service) DeleteRig(user *model.User, id uint) error {
	var data model.Rig

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteRig(&data)
}
