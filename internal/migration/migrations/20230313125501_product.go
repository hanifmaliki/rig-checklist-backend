package migrations

import (
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var mig20230313125501 = gormigrate.Migration{
	ID: "20230313125501",
	Migrate: func(tx *gorm.DB) error {
		type Product struct {
			model.GormAudit
			Slug                    string `json:"slug" gorm:"uniqueIndex:uidx_products_slug;size:191"`
			Name                    string `json:"name"`
			Desc                    string `json:"desc"`
			BannerImage             string `json:"banner_image"`
			BannerImageStyle        string `json:"banner_image_style"`
			BannerBackground        string `json:"banner_background"`
			HighlightedFeatureDesc  string `json:"highlighted_feature_desc"`
			HighlightedFeatureImage string `json:"highlighted_feature_image"`
			IsActive                bool   `json:"is_active"`
			NotArchived             bool   `json:"not_archived" gorm:"->;type:boolean GENERATED ALWAYS AS (IF(deleted_at IS NULL, 1, NULL)) VIRTUAL;default:(-);uniqueIndex:uidx_products_slug"`
		}
		return tx.AutoMigrate(&Product{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("products")
	},
}
