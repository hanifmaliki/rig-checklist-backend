package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type CaseStudyRepository interface {
	CreateCaseStudy(data *model.CaseStudy) (*model.CaseStudy, error)
	ReadCaseStudies(conds map[string]interface{}) ([]*model.CaseStudy, error)
	ReadCaseStudy(conds map[string]interface{}) (*model.CaseStudy, error)
	UpdateCaseStudy(data *model.CaseStudy) (*model.CaseStudy, error)
	DeleteCaseStudy(data *model.CaseStudy) error
}

func (r *repository) CreateCaseStudy(data *model.CaseStudy) (*model.CaseStudy, error) {
	db := persistence.Client().Minerva
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

func (r *repository) ReadCaseStudy(conds map[string]interface{}) (*model.CaseStudy, error) {
	db := persistence.Client().Minerva

	var data model.CaseStudy
	result := db.Preload("ImpactList", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_no, title")
	}).Preload("SolutionList", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product").Order("order_no")
	}).Where(conds).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadCaseStudies(conds map[string]interface{}) ([]*model.CaseStudy, error) {
	db := persistence.Client().Minerva

	var data []*model.CaseStudy
	result := db.Where(conds).Order("updated_at desc").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateCaseStudy(caseStudyInput *model.CaseStudy) (*model.CaseStudy, error) {
	db := persistence.Client().Minerva
	var caseStudy model.CaseStudy

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&caseStudy, caseStudyInput.ID)
		if result.Error != nil {
			return result.Error
		}

		var caseStudyImpact []*model.CaseStudyImpact
		result = db.Where("case_study_id = ?", caseStudyInput.ID).Find(&caseStudyImpact)
		if result.Error != nil {
			return result.Error
		}
		if len(caseStudyImpact) > 0 {
			result = tx.Model(&caseStudyImpact).UpdateColumn("deleted_by", caseStudyInput.UpdatedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&caseStudyImpact)
			if result.Error != nil {
				return result.Error
			}
		}

		var caseStudySolution []*model.CaseStudySolution
		result = db.Where("case_study_id = ?", caseStudyInput.ID).Find(&caseStudySolution)
		if result.Error != nil {
			return result.Error
		}
		if len(caseStudySolution) > 0 {
			result = tx.Model(&caseStudySolution).UpdateColumn("deleted_by", caseStudyInput.UpdatedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&caseStudySolution)
			if result.Error != nil {
				return result.Error
			}
		}

		caseStudy.Slug = caseStudyInput.Slug
		caseStudy.CompanyName = caseStudyInput.CompanyName
		caseStudy.Desc = caseStudyInput.Desc
		caseStudy.PersonName = caseStudyInput.PersonName
		caseStudy.PersonPosition = caseStudyInput.PersonPosition
		caseStudy.Logo = caseStudyInput.Logo
		caseStudy.Banner = caseStudyInput.Banner
		caseStudy.ImpactDesc = caseStudyInput.ImpactDesc
		caseStudy.IsActive = caseStudyInput.IsActive
		caseStudy.ImpactList = caseStudyInput.ImpactList
		caseStudy.SolutionList = caseStudyInput.SolutionList
		caseStudy.UpdatedBy = caseStudyInput.UpdatedBy

		result = tx.Save(&caseStudy)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &caseStudy, err
	}
}

func (r *repository) DeleteCaseStudy(caseStudyInput *model.CaseStudy) error {
	db := persistence.Client().Minerva
	var caseStudy model.CaseStudy

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&caseStudy, caseStudyInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&caseStudy).UpdateColumn("deleted_by", caseStudyInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&caseStudy)
		if result.Error != nil {
			return result.Error
		}

		var caseStudyImpact []*model.CaseStudyImpact
		result = db.Where("case_study_id = ?", caseStudyInput.ID).Find(&caseStudyImpact)
		if result.Error != nil {
			return result.Error
		}
		if len(caseStudyImpact) > 0 {
			result = tx.Model(&caseStudyImpact).UpdateColumn("deleted_by", caseStudyInput.DeletedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&caseStudyImpact)
			if result.Error != nil {
				return result.Error
			}
		}

		var caseStudySolution []*model.CaseStudySolution
		result = db.Where("case_study_id = ?", caseStudyInput.ID).Find(&caseStudySolution)
		if result.Error != nil {
			return result.Error
		}
		if len(caseStudySolution) > 0 {
			result = tx.Model(&caseStudySolution).UpdateColumn("deleted_by", caseStudyInput.DeletedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&caseStudySolution)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}
