package service

import (
	"sync"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/repository"
)

var (
	initOnce         sync.Once
	singletonService Service
)

type Service interface {
	AuthService
	HomeContentService
	ProductService
	ProductMenuService
	CaseStudyService
	CaseStudyMenuService
	FooterService
	FileService
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func init() {
	initOnce.Do(func() {
		singletonService = NewService(repository.NewRepository())
	})
}

func Instance() Service {
	return singletonService
}
