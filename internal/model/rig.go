package model

type Rig struct {
	GormAudit
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
