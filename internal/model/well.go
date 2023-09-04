package model

type Well struct {
	GormAudit
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
