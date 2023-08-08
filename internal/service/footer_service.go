package service

import (
	"strings"
	"time"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type FooterService interface {
	ReadFooter() (*model.Footer, error)
	UpdateFooter(user *model.User, footer *model.Footer) (*model.Footer, error)
}

func (s *service) ReadFooter() (*model.Footer, error) {
	return s.repository.ReadFooter()
}

func (s *service) UpdateFooter(user *model.User, footer *model.Footer) (*model.Footer, error) {
	footer.UpdatedBy = user.Email
	footer.Address = strings.TrimSpace(footer.Address)

	for idx := range footer.SitemapList {
		footer.SitemapList[idx].ID = 0
		footer.SitemapList[idx].CreatedAt = time.Time{}
		footer.SitemapList[idx].UpdatedAt = time.Time{}
		footer.SitemapList[idx].CreatedBy = user.Email
		footer.SitemapList[idx].UpdatedBy = user.Email
	}

	return s.repository.UpdateFooter(footer)
}
