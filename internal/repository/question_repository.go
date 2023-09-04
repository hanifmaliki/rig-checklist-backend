package repository

import (
	"sort"

	"gorm.io/gorm"

	"github.com/hanifmaliki/rig-checklist-backend/internal/helper"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/persistence"
)

type QuestionRepository interface {
	CreateQuestion(data *model.Question) (*model.Question, error)
	ReadQuestions(conds map[string]interface{}, sortBy string, analysis bool, limit int) ([]*model.Question, error)
	ReadQuestion(id uint, analysis bool) (*model.Question, error)
}

func (r *repository) CreateQuestion(data *model.Question) (*model.Question, error) {
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

func (r *repository) ReadQuestion(id uint, analysis bool) (*model.Question, error) {
	db := persistence.Client().RigChecklist
	var data model.Question

	result := db.Preload("SubSection").Preload("Classification")
	if analysis {
		result = result.Preload("Answers")
	}
	result = result.First(&data, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if analysis {
		helper.QuestionAnalysis(&data)
	}

	return &data, nil
}

func (r *repository) ReadQuestions(conds map[string]interface{}, sortBy string, analysis bool, limit int) ([]*model.Question, error) {
	db := persistence.Client().RigChecklist
	var data []*model.Question

	result := db.Preload("SubSection").Preload("Classification")
	if analysis {
		result = result.Preload("Answers")
	}
	result = result.Where(conds).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	if analysis {
		for _, question := range data {
			helper.QuestionAnalysis(question)
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i].TotalAnswers > data[j].TotalAnswers
		})
	}

	if limit != 0 && len(data) > limit {
		data = data[:limit]
	}

	return data, nil
}
