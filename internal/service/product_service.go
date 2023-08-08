package service

import (
	"strings"
	"time"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type ProductService interface {
	CreateProduct(user *model.User, product *model.Product) (*model.Product, error)
	ReadProducts(conds map[string]interface{}) ([]*model.Product, error)
	ReadProduct(conds map[string]interface{}) (*model.Product, error)
	UpdateProduct(user *model.User, id uint, product *model.Product) (*model.Product, error)
	DeleteProduct(user *model.User, id uint) error
}

func (s *service) CreateProduct(user *model.User, product *model.Product) (*model.Product, error) {
	product.ID = 0
	product.CreatedBy = user.Email
	product.UpdatedBy = user.Email
	product.Desc = strings.TrimSpace(product.Desc)
	product.HighlightedFeatureDesc = strings.TrimSpace(product.HighlightedFeatureDesc)

	for idx := range product.HighlightedFeatureList {
		product.HighlightedFeatureList[idx].ID = 0
		product.HighlightedFeatureList[idx].CreatedBy = user.Email
		product.HighlightedFeatureList[idx].UpdatedBy = user.Email
		product.HighlightedFeatureList[idx].Desc = strings.TrimSpace(product.HighlightedFeatureList[idx].Desc)
	}

	return s.repository.CreateProduct(product)
}

func (s *service) ReadProducts(conds map[string]interface{}) ([]*model.Product, error) {
	return s.repository.ReadProducts(conds)
}

func (s *service) ReadProduct(conds map[string]interface{}) (*model.Product, error) {
	return s.repository.ReadProduct(conds)
}

func (s *service) UpdateProduct(user *model.User, id uint, product *model.Product) (*model.Product, error) {
	product.ID = id
	product.UpdatedBy = user.Email
	product.Desc = strings.TrimSpace(product.Desc)
	product.HighlightedFeatureDesc = strings.TrimSpace(product.HighlightedFeatureDesc)

	for idx := range product.HighlightedFeatureList {
		product.HighlightedFeatureList[idx].ID = 0
		product.HighlightedFeatureList[idx].CreatedAt = time.Time{}
		product.HighlightedFeatureList[idx].UpdatedAt = time.Time{}
		product.HighlightedFeatureList[idx].CreatedBy = user.Email
		product.HighlightedFeatureList[idx].UpdatedBy = user.Email
		product.HighlightedFeatureList[idx].Desc = strings.TrimSpace(product.HighlightedFeatureList[idx].Desc)
	}

	return s.repository.UpdateProduct(product)
}

func (s *service) DeleteProduct(user *model.User, id uint) error {
	var product model.Product

	product.ID = id
	product.DeletedBy = user.Email
	return s.repository.DeleteProduct(&product)
}
