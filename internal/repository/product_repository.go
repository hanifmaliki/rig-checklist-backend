package repository

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(data *model.Product) (*model.Product, error)
	ReadProducts(conds map[string]interface{}) ([]*model.Product, error)
	ReadProduct(conds map[string]interface{}) (*model.Product, error)
	UpdateProduct(data *model.Product) (*model.Product, error)
	DeleteProduct(data *model.Product) error
}

func (r *repository) CreateProduct(data *model.Product) (*model.Product, error) {
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

func (r *repository) ReadProduct(conds map[string]interface{}) (*model.Product, error) {
	db := persistence.Client().Minerva

	var data model.Product
	result := db.Preload("HighlightedFeatureList", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_no, name")
	}).Where(conds).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *repository) ReadProducts(conds map[string]interface{}) ([]*model.Product, error) {
	db := persistence.Client().Minerva

	var data []*model.Product
	result := db.Where(conds).Order("updated_at desc").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *repository) UpdateProduct(productInput *model.Product) (*model.Product, error) {
	db := persistence.Client().Minerva
	var product model.Product

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&product, productInput.ID)
		if result.Error != nil {
			return result.Error
		}

		var productHighlightedFeature []*model.ProductHighlightedFeature
		result = db.Where("product_id = ?", productInput.ID).Find(&productHighlightedFeature)
		if result.Error != nil {
			return result.Error
		}
		if len(productHighlightedFeature) > 0 {
			result = tx.Model(&productHighlightedFeature).UpdateColumn("deleted_by", productInput.UpdatedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&productHighlightedFeature)
			if result.Error != nil {
				return result.Error
			}
		}

		product.Slug = productInput.Slug
		product.Name = productInput.Name
		product.Desc = productInput.Desc
		product.BannerImage = productInput.BannerImage
		product.BannerImageStyle = productInput.BannerImageStyle
		product.BannerBackground = productInput.BannerBackground
		product.HighlightedFeatureDesc = productInput.HighlightedFeatureDesc
		product.HighlightedFeatureImage = productInput.HighlightedFeatureImage
		product.IsActive = productInput.IsActive
		product.HighlightedFeatureList = productInput.HighlightedFeatureList
		product.UpdatedBy = productInput.UpdatedBy

		result = tx.Save(&product)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &product, err
	}
}

func (r *repository) DeleteProduct(productInput *model.Product) error {
	db := persistence.Client().Minerva
	var product model.Product

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&product, productInput.ID)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&product).UpdateColumn("deleted_by", productInput.DeletedBy)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&product)
		if result.Error != nil {
			return result.Error
		}

		var productHighlightedFeature []*model.ProductHighlightedFeature
		result = db.Where("product_id = ?", productInput.ID).Find(&productHighlightedFeature)
		if result.Error != nil {
			return result.Error
		}
		if len(productHighlightedFeature) > 0 {
			result = tx.Model(&productHighlightedFeature).UpdateColumn("deleted_by", productInput.DeletedBy)
			if result.Error != nil {
				return result.Error
			}
			result = tx.Delete(&productHighlightedFeature)
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
