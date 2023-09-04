package model

type Location struct {
	GormAudit
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
