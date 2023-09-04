package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	CreateAnswer(data *model.Answer) (*model.Answer, error)
	ReadAnswers(conds map[string]interface{}, sortBy string) ([]*model.Answer, error)
	ReadAnswer(id uint) (*model.Answer, error)
	UpdateAnswer(data *model.Answer) (*model.Answer, error)
	DeleteAnswer(data *model.Answer) error
}

func (r *repository) CreateAnswer(data *model.Answer) (*model.Answer, error) {
	db := persistence.Client().RigChecklist
	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(data)
		if err := result.Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return data, err
	}
}

func (r *repository) ReadAnswer(id uint) (*model.Answer, error) {
	db := persistence.Client().RigChecklist
	var data model.Answer

	result := db.Preload("Question").Preload("Photos").First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadAnswers(conds map[string]interface{}, sortBy string) ([]*model.Answer, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Answer

	if sortBy == "" {
		sortBy = "updated_at desc"
	}
	result := db.Preload("Question").Where(conds).Order(sortBy).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateAnswer(answerInput *model.Answer) (*model.Answer, error) {
	db := persistence.Client().RigChecklist
	var answer model.Answer

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&answer, answerInput.ID)
		if result.Error != nil {
			return result.Error
		}

		answer.FormID = answerInput.FormID
		answer.QuestionID = answerInput.QuestionID
		answer.IsExist = answerInput.IsExist
		answer.IsGood = answerInput.IsGood
		answer.NA = answerInput.NA
		answer.Note = answerInput.Note
		answer.UpdatedBy = answerInput.UpdatedBy

		result = tx.Save(&answer)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &answer, err
	}
}

func (r *repository) DeleteAnswer(answerInput *model.Answer) error {
	db := persistence.Client().RigChecklist
	var answer model.Answer

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&answer, answerInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&answer).UpdateColumn("deleted_by", answerInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&answer)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}
