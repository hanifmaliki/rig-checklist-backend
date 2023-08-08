package model

type Footer struct {
	GormAuditWithoutSoftDelete
	Address     string           `json:"address"`
	Phone       string           `json:"phone"`
	Email       string           `json:"email"`
	SitemapList []*FooterSitemap `json:"sitemap_list,omitempty" gorm:"-"`
}

type FooterSitemap struct {
	GormAudit
	Name    string `json:"name"`
	Url     string `json:"url"`
	OrderNo int    `json:"order_no"`
}
