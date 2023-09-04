package service

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
)

type FormService interface {
	CreateForm(user *model.User, data *model.Form) (*model.Form, error)
	ReadForms(conds map[string]interface{}, sortBy string) ([]*model.Form, error)
	ReadForm(id uint) (*model.Form, error)
	UpdateForm(user *model.User, id uint, data *model.Form) (*model.Form, error)
	DeleteForm(user *model.User, id uint) error

	StatusForm(user *model.User, id uint) error
}

func (s *service) CreateForm(user *model.User, data *model.Form) (*model.Form, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	questions, err := s.ReadQuestions(map[string]interface{}{}, "", false, 0)
	if err != nil {
		return nil, err
	}
	for _, question := range questions {
		data.Answers = append(data.Answers, &model.Answer{
			QuestionID: question.ID,
		})
	}

	return s.repository.CreateForm(data)
}

func (s *service) ReadForms(conds map[string]interface{}, sortBy string) ([]*model.Form, error) {
	return s.repository.ReadForms(conds, sortBy)
}

func (s *service) ReadForm(id uint) (*model.Form, error) {
	return s.repository.ReadForm(id)
}

func (s *service) UpdateForm(user *model.User, id uint, data *model.Form) (*model.Form, error) {
	data.ID = id
	data.UpdatedBy = user.Email

	return s.repository.UpdateForm(data)
}

func (s *service) DeleteForm(user *model.User, id uint) error {
	var data model.Form

	data.ID = id
	data.DeletedBy = user.Email
	return s.repository.DeleteForm(&data)
}
