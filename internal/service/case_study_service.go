package service

import (
	"strings"
	"time"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type CaseStudyService interface {
	CreateCaseStudy(user *model.User, caseStudy *model.CaseStudy) (*model.CaseStudy, error)
	ReadCaseStudies(conds map[string]interface{}) ([]*model.CaseStudy, error)
	ReadCaseStudy(conds map[string]interface{}) (*model.CaseStudy, error)
	UpdateCaseStudy(user *model.User, id uint, caseStudy *model.CaseStudy) (*model.CaseStudy, error)
	DeleteCaseStudy(user *model.User, id uint) error
}

func (s *service) CreateCaseStudy(user *model.User, caseStudy *model.CaseStudy) (*model.CaseStudy, error) {
	caseStudy.ID = 0
	caseStudy.CreatedBy = user.Email
	caseStudy.UpdatedBy = user.Email
	caseStudy.Desc = strings.TrimSpace(caseStudy.Desc)
	caseStudy.ImpactDesc = strings.TrimSpace(caseStudy.ImpactDesc)

	for idx := range caseStudy.ImpactList {
		caseStudy.ImpactList[idx].ID = 0
		caseStudy.ImpactList[idx].CreatedBy = user.Email
		caseStudy.ImpactList[idx].UpdatedBy = user.Email
		caseStudy.ImpactList[idx].Desc = strings.TrimSpace(caseStudy.ImpactList[idx].Desc)
	}

	for idx := range caseStudy.SolutionList {
		caseStudy.SolutionList[idx].ID = 0
		caseStudy.SolutionList[idx].CreatedBy = user.Email
		caseStudy.SolutionList[idx].UpdatedBy = user.Email
	}

	return s.repository.CreateCaseStudy(caseStudy)
}

func (s *service) ReadCaseStudies(conds map[string]interface{}) ([]*model.CaseStudy, error) {
	return s.repository.ReadCaseStudies(conds)
}

func (s *service) ReadCaseStudy(conds map[string]interface{}) (*model.CaseStudy, error) {
	return s.repository.ReadCaseStudy(conds)
}

func (s *service) UpdateCaseStudy(user *model.User, id uint, caseStudy *model.CaseStudy) (*model.CaseStudy, error) {
	caseStudy.ID = id
	caseStudy.UpdatedBy = user.Email
	caseStudy.Desc = strings.TrimSpace(caseStudy.Desc)
	caseStudy.ImpactDesc = strings.TrimSpace(caseStudy.ImpactDesc)

	for idx := range caseStudy.ImpactList {
		caseStudy.ImpactList[idx].ID = 0
		caseStudy.ImpactList[idx].CreatedAt = time.Time{}
		caseStudy.ImpactList[idx].UpdatedAt = time.Time{}
		caseStudy.ImpactList[idx].CreatedBy = user.Email
		caseStudy.ImpactList[idx].UpdatedBy = user.Email
		caseStudy.ImpactList[idx].Desc = strings.TrimSpace(caseStudy.ImpactList[idx].Desc)
	}

	for idx := range caseStudy.SolutionList {
		caseStudy.SolutionList[idx].ID = 0
		caseStudy.SolutionList[idx].CreatedAt = time.Time{}
		caseStudy.SolutionList[idx].UpdatedAt = time.Time{}
		caseStudy.SolutionList[idx].CreatedBy = user.Email
		caseStudy.SolutionList[idx].UpdatedBy = user.Email
	}

	return s.repository.UpdateCaseStudy(caseStudy)
}

func (s *service) DeleteCaseStudy(user *model.User, id uint) error {
	var caseStudy model.CaseStudy

	caseStudy.ID = id
	caseStudy.DeletedBy = user.Email
	return s.repository.DeleteCaseStudy(&caseStudy)
}
