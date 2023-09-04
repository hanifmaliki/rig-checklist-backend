package service

import (
	"sync"

	"github.com/hanifmaliki/rig-checklist-backend/internal/repository"
)

var (
	initOnce         sync.Once
	singletonService Service
)

type Service interface {
	UserService
	AuthService
	ActivityService
	FieldService
	LocationService
	RigService
	WellService
	QuestionService
	AnswerService
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
