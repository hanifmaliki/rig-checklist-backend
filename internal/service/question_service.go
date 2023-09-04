package service

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
)

type QuestionService interface {
	CreateQuestion(user *model.User, data *model.Question) (*model.Question, error)
	ReadQuestions(conds map[string]interface{}, sortBy string, analysis bool, limit int) ([]*model.Question, error)
	ReadQuestion(id uint, analysis bool) (*model.Question, error)
}

func (s *service) CreateQuestion(user *model.User, data *model.Question) (*model.Question, error) {
	data.ID = 0
	data.CreatedBy = user.Email
	data.UpdatedBy = user.Email

	return s.repository.CreateQuestion(data)
}

func (s *service) ReadQuestions(conds map[string]interface{}, sortBy string, analysis bool, limit int) ([]*model.Question, error) {
	return s.repository.ReadQuestions(conds, sortBy, analysis, limit)
}

func (s *service) ReadQuestion(id uint, analysis bool) (*model.Question, error) {
	return s.repository.ReadQuestion(id, analysis)
}
