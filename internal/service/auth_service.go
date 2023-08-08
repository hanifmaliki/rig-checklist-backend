package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type AuthService interface {
	Login(login *model.Login) (*jwt.Token, error)
}

func (s *service) Login(login *model.Login) (*jwt.Token, error) {
	user := &model.User{}
	if login.Username != os.Getenv("ADMIN_USERNAME") || login.Password != os.Getenv("ADMIN_PASSWORD") {
		success, err := helper.LoginPetros(login.Username, login.Password)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, errors.New("the username or password is incorrect")
		}
		userPetros, err := s.repository.ReadUserPetrosByUsername(login.Username)
		if err != nil {
			return nil, err
		}
		user.Username = userPetros.UserLogin
		user.Name = userPetros.DisplayName
		user.Email = userPetros.UserEmail
	} else {
		user.Username = login.Username
		user.Name = helper.UserDummy.Name
		user.Email = helper.UserDummy.Email
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token, nil
}
