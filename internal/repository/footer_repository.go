package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type FooterRepository interface {
	ReadFooter() (*model.Footer, error)
	UpdateFooter(footer *model.Footer) (*model.Footer, error)
}

func (r *repository) ReadFooter() (*model.Footer, error) {
	db := persistence.Client().Minerva
	var footer model.Footer
	result := db.Order("updated_at desc").First(&footer)
	if result.Error != nil {
		return nil, result.Error
	}

	var footerSitemaps []*model.FooterSitemap
	result = db.Order("order_no, name").Find(&footerSitemaps)
	if result.Error != nil {
		return nil, result.Error
	}
	footer.SitemapList = footerSitemaps

	return &footer, nil
}

func (r *repository) UpdateFooter(footerInput *model.Footer) (*model.Footer, error) {
	db := persistence.Client().Minerva
	footer := model.Footer{}

	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Order("updated_at desc").First(&footer)
		if result.Error != nil {
			return result.Error
		}
		footer.Address = footerInput.Address
		footer.Phone = footerInput.Phone
		footer.Email = footerInput.Email
		footer.UpdatedBy = footerInput.UpdatedBy
		result = tx.Save(&footer)
		if result.Error != nil {
			return result.Error
		}

		if len(footerInput.SitemapList) > 0 {
			err := deleteFooterSitemaps(db, tx, &footer.UpdatedBy)
			if err != nil {
				return err
			}

			result = tx.Create(&footerInput.SitemapList)
			if result.Error != nil {
				return result.Error
			}
			footer.SitemapList = footerInput.SitemapList
		} else {
			err := deleteFooterSitemaps(db, tx, &footer.UpdatedBy)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &footer, err
	}
}

func deleteFooterSitemaps(db *gorm.DB, tx *gorm.DB, deletedBy *string) error {
	var footerSitemaps []model.FooterSitemap
	result := db.Find(&footerSitemaps)
	if result.Error != nil {
		return result.Error
	}

	if len(footerSitemaps) > 0 {
		result = tx.Model(&footerSitemaps).UpdateColumn("deleted_by", deletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&footerSitemaps)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
