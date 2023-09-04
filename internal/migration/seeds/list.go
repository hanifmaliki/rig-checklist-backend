package seeds

import "gorm.io/gorm"

var Seeds = []func(db *gorm.DB) error{
	SeedUser,
	SeedPosition,
	SeedRig,
	SeedLocation,
	SeedWell,
	SeedField,
	SeedActivity,
	SeedSectionSubSection,
	SeedClassification,
	SeedQuestion,
	SeedQuestion,
	SeedFormDummy,
	SeedPhotoDummy,
	SeedSignDummy,
}
