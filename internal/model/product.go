package model

type Product struct {
	GormAudit
	Slug                    string                       `json:"slug" gorm:"uniqueIndex:uidx_products_slug;size:191"`
	Name                    string                       `json:"name"`
	Desc                    string                       `json:"desc"`
	BannerImage             string                       `json:"banner_image"`
	BannerImageStyle        string                       `json:"banner_image_style"`
	BannerBackground        string                       `json:"banner_background"`
	HighlightedFeatureDesc  string                       `json:"highlighted_feature_desc"`
	HighlightedFeatureImage string                       `json:"highlighted_feature_image"`
	IsActive                bool                         `json:"is_active"`
	HighlightedFeatureList  []*ProductHighlightedFeature `json:"highlighted_feature_list,omitempty"`
}

type ProductHighlightedFeature struct {
	GormAudit
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Image     string `json:"image"`
	OrderNo   int    `json:"order_no"`
}

type ProductMenu struct {
	GormAudit
	ProductID uint     `json:"product_id"`
	Type      string   `json:"type"`
	Desc      string   `json:"desc"`
	IsActive  bool     `json:"is_active"`
	OrderNo   int      `json:"order_no"`
	Product   *Product `json:"product,omitempty" gorm:"<-:false"`
}
