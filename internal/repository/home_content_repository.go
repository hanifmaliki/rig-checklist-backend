package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HomeContentRepository interface {
	UpdateHomeContents(homeContents []*model.HomeContent) ([]*model.HomeContent, error)
	ReadHomeContents(query *model.HomeContent) ([]*model.HomeContent, error)
	DeleteHomeContent(id uint) error
}

func (r *repository) UpdateHomeContents(homeContents []*model.HomeContent) ([]*model.HomeContent, error) {
	db := persistence.Client().Minerva
	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "section"}, {Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value", "is_json", "updated_by", "updated_at"}),
		}).Create(homeContents)
		if err := result.Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return homeContents, err
	}
}

func (r *repository) ReadHomeContents(query *model.HomeContent) ([]*model.HomeContent, error) {
	db := persistence.Client().Minerva
	var data []*model.HomeContent
	db.Where(query).Find(&data)
	return data, nil
}

func (r *repository) DeleteHomeContent(id uint) error {
	db := persistence.Client().Minerva
	var data model.HomeContent
	data.ID = id

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&data, id)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Delete(&data)
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
