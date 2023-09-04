package model

type Sign struct {
	GormAudit
	UserID uint   `json:"user_id"`
	Url    string `json:"url"`
}
