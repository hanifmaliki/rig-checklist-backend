package seeds

import "gorm.io/gorm"

var Seeds = []func(db *gorm.DB) error{
	SeedHomeContent,
	SeedProduct,
	SeedCaseStudy,
	SeedFooter,
	SeedFile,
}
