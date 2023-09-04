package model

type Activity struct {
	GormAudit
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
