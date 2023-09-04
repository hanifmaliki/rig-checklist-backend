package model

type Position struct {
	GormAuditWithoutSoftDelete
	Name string `json:"name"`
	Role string `json:"role"`
}
