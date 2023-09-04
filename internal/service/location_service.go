package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type LocationService interface {
	CreateLocation(user *model.User, data *model.Location) (*model.Location, error)
	ReadLocations(conds map[string]interface{}, sortBy string) ([]*model.Location, error)
	ReadLocation(id uint) (*model.Location, error)
	UpdateLocation(user *model.User, id uint, data *model.Location) (*model.Location, error)
	DeleteLocation(user *model.User, id uint) error
}

func (s *service) CreateLocation(user *model.User, data *model.Location) (*model.Location, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateLocation(data)
}

func (s *service) ReadLocations(conds map[string]interface{}, sortBy string) ([]*model.Location, error) {
	return s.repository.ReadLocations(conds, sortBy)
}

func (s *service) ReadLocation(id uint) (*model.Location, error) {
	return s.repository.ReadLocation(id)
}

func (s *service) UpdateLocation(user *model.User, id uint, data *model.Location) (*model.Location, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateLocation(data)
}

func (s *service) DeleteLocation(user *model.User, id uint) error {
	var data model.Location

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteLocation(&data)
}
