package model

type SubSection struct {
	GormAuditWithoutSoftDelete
	SectionID uint        `json:"section_id"`
	Name      string      `json:"name"`
	Section   *Section    `json:"section,omitempty"`
	Questions []*Question `json:"questions,omitempty"`
}
