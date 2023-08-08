package repository

import (
	"errors"
	"strconv"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type CaseStudyMenuRepository interface {
	CreateCaseStudyMenu(caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error)
	ReadCaseStudyMenus(conds map[string]interface{}, orderBy string) ([]*model.CaseStudyMenu, error)
	ReadCaseStudyMenu(id uint) (*model.CaseStudyMenu, error)
	UpdateCaseStudyMenu(caseStudyMenuInput *model.CaseStudyMenu) (*model.CaseStudyMenu, error)
	DeleteCaseStudyMenu(caseStudyMenu *model.CaseStudyMenu) error
}

func (r *repository) CreateCaseStudyMenu(caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error) {
	db := persistence.Client().Minerva

	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(caseStudyMenu)
		if err := result.Error; err != nil {
			return err
		}

		var caseStudy model.CaseStudy
		result = db.First(&caseStudy, caseStudyMenu.CaseStudyID)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return errors.New("case_study_id " + strconv.Itoa(int(caseStudyMenu.CaseStudyID)) + " not found")
			} else {
				return result.Error
			}
		}
		caseStudyMenu.CaseStudy = &caseStudy

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return caseStudyMenu, err
	}
}

func (r *repository) ReadCaseStudyMenu(id uint) (*model.CaseStudyMenu, error) {
	db := persistence.Client().Minerva

	var caseStudyMenu model.CaseStudyMenu
	result := db.Preload("CaseStudy").First(&caseStudyMenu, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &caseStudyMenu, nil
}

func (r *repository) ReadCaseStudyMenus(conds map[string]interface{}, orderBy string) ([]*model.CaseStudyMenu, error) {
	db := persistence.Client().Minerva

	var caseStudyMenus []*model.CaseStudyMenu
	if orderBy == "" {
		orderBy = "order_no"
	}
	db.Preload("CaseStudy").Where(conds).Order(orderBy).Find(&caseStudyMenus)

	return caseStudyMenus, nil
}

func (r *repository) UpdateCaseStudyMenu(caseStudyMenuInput *model.CaseStudyMenu) (*model.CaseStudyMenu, error) {
	db := persistence.Client().Minerva
	caseStudyMenu := model.CaseStudyMenu{}

	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.First(&caseStudyMenu, caseStudyMenuInput.ID)
		if result.Error != nil {
			return result.Error
		}
		caseStudyMenu.CaseStudyID = caseStudyMenuInput.CaseStudyID
		caseStudyMenu.Type = caseStudyMenuInput.Type
		caseStudyMenu.Desc = caseStudyMenuInput.Desc
		caseStudyMenu.IsActive = caseStudyMenuInput.IsActive
		caseStudyMenu.OrderNo = caseStudyMenuInput.OrderNo
		caseStudyMenu.UpdatedBy = caseStudyMenuInput.UpdatedBy
		result = tx.Save(&caseStudyMenu)
		if result.Error != nil {
			return result.Error
		}

		var caseStudy model.CaseStudy
		result = db.First(&caseStudy, caseStudyMenu.CaseStudyID)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return errors.New("case_study_id " + strconv.Itoa(int(caseStudyMenu.CaseStudyID)) + " not found")
			} else {
				return result.Error
			}
		}
		caseStudyMenu.CaseStudy = &caseStudy

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &caseStudyMenu, err
	}
}

func (r *repository) DeleteCaseStudyMenu(caseStudyMenu *model.CaseStudyMenu) error {
	db := persistence.Client().Minerva
	data := model.CaseStudyMenu{}

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&data, caseStudyMenu.ID)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Model(&data).UpdateColumn("deleted_by", caseStudyMenu.DeletedBy)
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
