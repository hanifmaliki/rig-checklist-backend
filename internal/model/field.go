package model

type Field struct {
	GormAudit
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
