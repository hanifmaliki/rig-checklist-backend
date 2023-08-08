package repository

import (
	"errors"
	"strconv"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/persistence"

	"gorm.io/gorm"
)

type ProductMenuRepository interface {
	CreateProductMenu(productMenu *model.ProductMenu) (*model.ProductMenu, error)
	ReadProductMenus(conds map[string]interface{}, orderBy string) ([]*model.ProductMenu, error)
	ReadProductMenu(id uint) (*model.ProductMenu, error)
	UpdateProductMenu(productMenuInput *model.ProductMenu) (*model.ProductMenu, error)
	DeleteProductMenu(productMenu *model.ProductMenu) error
}

func (r *repository) CreateProductMenu(productMenu *model.ProductMenu) (*model.ProductMenu, error) {
	db := persistence.Client().Minerva

	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(productMenu)
		if err := result.Error; err != nil {
			return err
		}

		var product model.Product
		result = db.First(&product, productMenu.ProductID)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return errors.New("product_id " + strconv.Itoa(int(productMenu.ProductID)) + " not found")
			} else {
				return result.Error
			}
		}
		productMenu.Product = &product

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return productMenu, err
	}
}

func (r *repository) ReadProductMenu(id uint) (*model.ProductMenu, error) {
	db := persistence.Client().Minerva

	var productMenu model.ProductMenu
	result := db.Preload("Product").First(&productMenu, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &productMenu, nil
}

func (r *repository) ReadProductMenus(conds map[string]interface{}, orderBy string) ([]*model.ProductMenu, error) {
	db := persistence.Client().Minerva

	var productMenus []*model.ProductMenu
	if orderBy == "" {
		orderBy = "order_no"
	}
	db.Preload("Product").Where(conds).Order(orderBy).Find(&productMenus)

	return productMenus, nil
}

func (r *repository) UpdateProductMenu(productMenuInput *model.ProductMenu) (*model.ProductMenu, error) {
	db := persistence.Client().Minerva
	productMenu := model.ProductMenu{}

	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.First(&productMenu, productMenuInput.ID)
		if result.Error != nil {
			return result.Error
		}
		productMenu.ProductID = productMenuInput.ProductID
		productMenu.Type = productMenuInput.Type
		productMenu.Desc = productMenuInput.Desc
		productMenu.IsActive = productMenuInput.IsActive
		productMenu.OrderNo = productMenuInput.OrderNo
		productMenu.UpdatedBy = productMenuInput.UpdatedBy
		result = tx.Save(&productMenu)
		if result.Error != nil {
			return result.Error
		}

		var product model.Product
		result = db.First(&product, productMenu.ProductID)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return errors.New("product_id " + strconv.Itoa(int(productMenu.ProductID)) + " not found")
			} else {
				return result.Error
			}
		}
		productMenu.Product = &product

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &productMenu, err
	}
}

func (r *repository) DeleteProductMenu(productMenu *model.ProductMenu) error {
	db := persistence.Client().Minerva
	data := model.ProductMenu{}

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.First(&data, productMenu.ID)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Model(&data).UpdateColumn("deleted_by", productMenu.DeletedBy)
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
