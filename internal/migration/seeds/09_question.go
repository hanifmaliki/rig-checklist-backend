package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gorm.io/gorm"
)

func SeedQuestion(db *gorm.DB) error {
	seeds := []*model.Question{
		{
			No:               "1",
			SubNo:            "",
			Question:         "Program Sumuran (Well Service) yang sudah ditanda tangani oleh pejabat terkait:",
			ClassificationID: 2,
		},
		{
			No:               "1",
			SubNo:            "a",
			Question:         "Program Sumuran (Well Service) yang sudah ditanda tangani oleh pejabat terkait:",
			ClassificationID: 2,
		},
	}

	for idx := range seeds {
		seeds[idx].CreatedBy = helper.UserDummy.Email
		seeds[idx].UpdatedBy = helper.UserDummy.Email
	}

	result := db.Create(&seeds)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
