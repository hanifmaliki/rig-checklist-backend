package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type CaseStudyMenuService interface {
	CreateCaseStudyMenu(user *model.User, caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error)
	ReadCaseStudyMenus(conds map[string]interface{}, orderBy string) ([]*model.CaseStudyMenu, error)
	ReadCaseStudyMenu(id uint) (*model.CaseStudyMenu, error)
	UpdateCaseStudyMenu(user *model.User, id uint, caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error)
	DeleteCaseStudyMenu(user *model.User, id uint) error
}

func (s *service) CreateCaseStudyMenu(user *model.User, caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error) {
	caseStudyMenu.ID = 0
	caseStudyMenu.CreatedBy = user.Email
	caseStudyMenu.UpdatedBy = user.Email

	return s.repository.CreateCaseStudyMenu(caseStudyMenu)
}

func (s *service) ReadCaseStudyMenus(conds map[string]interface{}, orderBy string) ([]*model.CaseStudyMenu, error) {
	return s.repository.ReadCaseStudyMenus(conds, orderBy)
}

func (s *service) ReadCaseStudyMenu(id uint) (*model.CaseStudyMenu, error) {
	return s.repository.ReadCaseStudyMenu(id)
}

func (s *service) UpdateCaseStudyMenu(user *model.User, id uint, caseStudyMenu *model.CaseStudyMenu) (*model.CaseStudyMenu, error) {
	caseStudyMenu.ID = id
	caseStudyMenu.UpdatedBy = user.Email
	return s.repository.UpdateCaseStudyMenu(caseStudyMenu)
}

func (s *service) DeleteCaseStudyMenu(user *model.User, id uint) error {
	var caseStudyMenu model.CaseStudyMenu

	caseStudyMenu.ID = id
	caseStudyMenu.DeletedBy = user.Email
	return s.repository.DeleteCaseStudyMenu(&caseStudyMenu)
}
