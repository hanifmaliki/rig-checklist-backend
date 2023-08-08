package service

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

type ProductMenuService interface {
	CreateProductMenu(user *model.User, productMenu *model.ProductMenu) (*model.ProductMenu, error)
	ReadProductMenus(conds map[string]interface{}, orderBy string) ([]*model.ProductMenu, error)
	ReadProductMenu(id uint) (*model.ProductMenu, error)
	UpdateProductMenu(user *model.User, id uint, productMenu *model.ProductMenu) (*model.ProductMenu, error)
	DeleteProductMenu(user *model.User, id uint) error
}

func (s *service) CreateProductMenu(user *model.User, productMenu *model.ProductMenu) (*model.ProductMenu, error) {
	productMenu.ID = 0
	productMenu.CreatedBy = user.Email
	productMenu.UpdatedBy = user.Email

	return s.repository.CreateProductMenu(productMenu)
}

func (s *service) ReadProductMenus(conds map[string]interface{}, orderBy string) ([]*model.ProductMenu, error) {
	return s.repository.ReadProductMenus(conds, orderBy)
}

func (s *service) ReadProductMenu(id uint) (*model.ProductMenu, error) {
	return s.repository.ReadProductMenu(id)
}

func (s *service) UpdateProductMenu(user *model.User, id uint, productMenu *model.ProductMenu) (*model.ProductMenu, error) {
	productMenu.ID = id
	productMenu.UpdatedBy = user.Email
	return s.repository.UpdateProductMenu(productMenu)
}

func (s *service) DeleteProductMenu(user *model.User, id uint) error {
	var productMenu model.ProductMenu

	productMenu.ID = id
	productMenu.DeletedBy = user.Email
	return s.repository.DeleteProductMenu(&productMenu)
}
