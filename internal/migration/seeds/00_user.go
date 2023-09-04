package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	seeds := []*model.User{
		{
			Name:       helper.UserDummy.Name,
			Username:   helper.UserDummy.Username,
			Email:      helper.UserDummy.Email,
			Password:   "Test123@",
			PositionID: 1,
			FieldID:    1,
			IsAdmin:    true,
			IsActive:   true,
		},
	}

	for _, seed := range seeds {
		_, err := service.Instance().CreateUser(helper.UserDummy, seed)
		if err != nil {
			return err
		}
	}

	return nil
}
