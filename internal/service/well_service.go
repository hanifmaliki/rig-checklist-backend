package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type WellService interface {
	CreateWell(user *model.User, data *model.Well) (*model.Well, error)
	ReadWells(conds map[string]interface{}, sortBy string) ([]*model.Well, error)
	ReadWell(id uint) (*model.Well, error)
	UpdateWell(user *model.User, id uint, data *model.Well) (*model.Well, error)
	DeleteWell(user *model.User, id uint) error
}

func (s *service) CreateWell(user *model.User, data *model.Well) (*model.Well, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateWell(data)
}

func (s *service) ReadWells(conds map[string]interface{}, sortBy string) ([]*model.Well, error) {
	return s.repository.ReadWells(conds, sortBy)
}

func (s *service) ReadWell(id uint) (*model.Well, error) {
	return s.repository.ReadWell(id)
}

func (s *service) UpdateWell(user *model.User, id uint, data *model.Well) (*model.Well, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateWell(data)
}

func (s *service) DeleteWell(user *model.User, id uint) error {
	var data model.Well

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteWell(&data)
}
