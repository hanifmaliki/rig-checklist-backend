package model

type Section struct {
	GormAuditWithoutSoftDelete
	Name        string        `json:"name"`
	SubSections []*SubSection `json:"sub_sections,omitempty"`
	Questions   []*Question   `json:"questions,omitempty"`
}
