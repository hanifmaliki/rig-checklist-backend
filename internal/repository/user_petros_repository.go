package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"
)

type UserPetrosRepository interface {
	ReadUserPetrosByUsername(username string) (*model.WpUser, error)
}

func (r *repository) ReadUserPetrosByUsername(username string) (*model.WpUser, error) {
	db := persistence.Client().Petros

	var userPetros model.WpUser
	result := db.Where(&model.WpUser{UserLogin: username}).First(&userPetros)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userPetros, nil
}
