package seeds

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gorm.io/gorm"
)

func SeedFooter(db *gorm.DB) error {
	footerSeed := model.Footer{
		Address: "Old Mineral House, Level 3, 2 Edward Street,<br/>Brisbane City Queensland 4000<br/>Australia",
		Phone:   "+61730510971",
		Email:   "info@petros-solutions.com",
		SitemapList: []*model.FooterSitemap{
			{
				Name:    "About Us",
				Url:     "https://petros-solutions.com/about-us/",
				OrderNo: 1,
			},
			{
				Name:    "Services",
				Url:     "https://petros-solutions.com/services/",
				OrderNo: 2,
			},
			{
				Name:    "Portfolio",
				Url:     "https://petros-solutions.com/portfolio/",
				OrderNo: 3,
			},
			{
				Name:    "Contact Us",
				Url:     "https://petros-solutions.com/contact-us/",
				OrderNo: 4,
			},
		},
	}

	footerSeed.CreatedBy = helper.UserDummy.Email
	footerSeed.UpdatedBy = helper.UserDummy.Email
	for idx := range footerSeed.SitemapList {
		footerSeed.SitemapList[idx].CreatedBy = helper.UserDummy.Email
		footerSeed.SitemapList[idx].UpdatedBy = helper.UserDummy.Email
	}

	result := db.Create(&footerSeed)
	if result.Error != nil {
		return result.Error
	}

	result = db.Create(&footerSeed.SitemapList)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
