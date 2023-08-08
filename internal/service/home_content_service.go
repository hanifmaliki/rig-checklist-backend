package service

import (
	"time"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type HomeContentService interface {
	UpdateHomeContents(user *model.User, homeContents []*model.HomeContent) ([]*model.HomeContent, error)
	ReadHomeContents(query *model.HomeContent) ([]*model.HomeContent, error)
	DeleteHomeContent(id uint) error
}

func (s *service) UpdateHomeContents(user *model.User, homeContents []*model.HomeContent) ([]*model.HomeContent, error) {
	for idx := range homeContents {
		homeContents[idx].CreatedAt = time.Time{}
		homeContents[idx].UpdatedAt = time.Time{}
		homeContents[idx].CreatedBy = user.Email
		homeContents[idx].UpdatedBy = user.Email
	}

	return s.repository.UpdateHomeContents(homeContents)
}

func (s *service) ReadHomeContents(query *model.HomeContent) ([]*model.HomeContent, error) {
	return s.repository.ReadHomeContents(query)
}

func (s *service) DeleteHomeContent(id uint) error {
	return s.repository.DeleteHomeContent(id)
}
