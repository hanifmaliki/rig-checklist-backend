package service

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type AuthService interface {
	Login(login *model.Login) (*jwt.Token, error)
	ChangePassword(user *model.User, oldPassword string, newPassword string) error

	// Forgot Password
	SendOTP(email string) error
	VerifyOTP(email string, otp string) (*jwt.Token, error)
	ResetPassword(user *model.User, newPassword string) error
}

func (s *service) Login(login *model.Login) (*jwt.Token, error) {
	conds := map[string]interface{}{
		"username": strings.ToLower(login.Username),
	}
	user, err := s.ReadUser(conds)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token, nil
}

func (s *service) ChangePassword(user *model.User, oldPassword string, newPassword string) error {
	conds := map[string]interface{}{
		"username": strings.ToLower(user.Username),
	}
	user, err := s.ReadUser(conds)
	if err != nil {
		return err
	}

	err = helper.ComparePassword(user.Password, oldPassword)
	if err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	_, err = s.UpdateUser(user, user.ID, user)
	return err
}

func (s *service) SendOTP(email string) error {
	return nil
}

func (s *service) VerifyOTP(email string, otp string) (*jwt.Token, error) {
	return nil, nil
}

func (s *service) ResetPassword(user *model.User, newPassword string) error {
	conds := map[string]interface{}{
		"username": strings.ToLower(user.Username),
	}
	user, err := s.ReadUser(conds)
	if err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	_, err = s.UpdateUser(user, user.ID, user)
	return err
}
