package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type AnswerService interface {
	CreateAnswer(user *model.User, data *model.Answer) (*model.Answer, error)
	ReadAnswers(conds map[string]interface{}, sortBy string) ([]*model.Answer, error)
	ReadAnswer(id uint) (*model.Answer, error)
	UpdateAnswer(user *model.User, id uint, data *model.Answer) (*model.Answer, error)
	DeleteAnswer(user *model.User, id uint) error
}

func (s *service) CreateAnswer(user *model.User, data *model.Answer) (*model.Answer, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateAnswer(data)
}

func (s *service) ReadAnswers(conds map[string]interface{}, sortBy string) ([]*model.Answer, error) {
	return s.repository.ReadAnswers(conds, sortBy)
}

func (s *service) ReadAnswer(id uint) (*model.Answer, error) {
	return s.repository.ReadAnswer(id)
}

func (s *service) UpdateAnswer(user *model.User, id uint, data *model.Answer) (*model.Answer, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateAnswer(data)
}

func (s *service) DeleteAnswer(user *model.User, id uint) error {
	var data model.Answer

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteAnswer(&data)
}
