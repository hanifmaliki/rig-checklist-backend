package model

type Photo struct {
	GormAudit
	AnswerID uint   `json:"answer_id"`
	Url      string `json:"url"`
}
