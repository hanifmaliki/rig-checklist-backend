package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type ActivityService interface {
	CreateActivity(user *model.User, data *model.Activity) (*model.Activity, error)
	ReadActivities(conds map[string]interface{}, sortBy string) ([]*model.Activity, error)
	ReadActivity(id uint) (*model.Activity, error)
	UpdateActivity(user *model.User, id uint, data *model.Activity) (*model.Activity, error)
	DeleteActivity(user *model.User, id uint) error
}

func (s *service) CreateActivity(user *model.User, data *model.Activity) (*model.Activity, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateActivity(data)
}

func (s *service) ReadActivities(conds map[string]interface{}, sortBy string) ([]*model.Activity, error) {
	return s.repository.ReadActivities(conds, sortBy)
}

func (s *service) ReadActivity(id uint) (*model.Activity, error) {
	return s.repository.ReadActivity(id)
}

func (s *service) UpdateActivity(user *model.User, id uint, data *model.Activity) (*model.Activity, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateActivity(data)
}

func (s *service) DeleteActivity(user *model.User, id uint) error {
	var data model.Activity

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteActivity(&data)
}
