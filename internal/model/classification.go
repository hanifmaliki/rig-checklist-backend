package model

type Classification struct {
	GormAuditWithoutSoftDelete
	Name string `json:"name"`
}
