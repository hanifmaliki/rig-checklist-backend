package model

type HomeContent struct {
	GormAuditWithoutSoftDelete
	Section string `json:"section" gorm:"uniqueIndex:uidx_home_contents;size:191"`
	Key     string `json:"key" gorm:"uniqueIndex:uidx_home_contents;size:191"`
	Value   string `json:"value"`
	IsJson  bool   `json:"is_json"`
}
