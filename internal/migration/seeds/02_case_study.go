package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"

	"gorm.io/gorm"
)

func SeedCaseStudy(db *gorm.DB) error {
	caseStudySeeds := []model.CaseStudy{
		{
			Slug:           "cs-petrosea",
			CompanyName:    "PT Petrosea Tbk",
			Desc:           `"Digitalisation is the catalyst that is supporting mining operations to become ‘smarter’ by leveraging digital tools and optimal processes that make operations instrumented, interconnected and intelligent."`,
			PersonName:     "Petrosea",
			PersonPosition: "Mining Directors",
			Logo:           "case-study-1-logo.webp",
			Banner:         "case-study-1-banner.webm",
			ImpactDesc:     "Centralized command and control from our ROC enables proactive management, monitoring and coordination while maintaining dependability",
			IsActive:       true,
			ImpactList: []*model.CaseStudyImpact{
				{
					Title:    "15%",
					SubTitle: "",
					Desc:     "Increased Labor Efficiency without reducing the work quality",
					Image:    "case-study-1-impact-1.svg",
					OrderNo:  1,
				},
				{
					Title:    "57 for 5",
					SubTitle: "",
					Desc:     "Reduce the number of people required for projects",
					Image:    "case-study-1-impact-2.svg",
					OrderNo:  2,
				},
				{
					Title:    "3%",
					SubTitle: "",
					Desc:     "Increased major fleet’s productivity due to interconnected digital products",
					Image:    "case-study-1-impact-3.svg",
					OrderNo:  3,
				},
			},
			SolutionList: []*model.CaseStudySolution{
				{
					ProductID: 7,
					OrderNo:   1,
				},
				{
					ProductID: 3,
					OrderNo:   1,
				},
			},
		},
	}

	caseStudyMenuSeeds := map[string]model.CaseStudyMenu{
		"PT Petrosea Tbk": {
			Type:     "",
			Desc:     "Engineering Services Company",
			IsActive: true,
			OrderNo:  1,
		},
	}

	for _, caseStudySeed := range caseStudySeeds {
		caseStudyResult, err := service.Instance().CreateCaseStudy(helper.UserDummy, &caseStudySeed)
		if err != nil {
			return err
		}

		caseStudyMenuSeed := caseStudyMenuSeeds[caseStudyResult.CompanyName]
		caseStudyMenuSeed.CaseStudyID = caseStudyResult.ID

		_, err = service.Instance().CreateCaseStudyMenu(helper.UserDummy, &caseStudyMenuSeed)
		if err != nil {
			return err
		}
	}

	return nil
}
