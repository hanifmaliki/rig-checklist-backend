package seeds

import (
	"github.com/hanifmaliki/rig-checklist-backend/internal/helper"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/service"

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
