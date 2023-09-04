package model

type Question struct {
	GormAuditWithoutSoftDelete
	SectionID         uint            `json:"section_id"`
	SubSectionID      uint            `json:"sub_section_id"`
	No                string          `json:"no"`
	SubNo             string          `json:"sub_no"`
	Question          string          `json:"question"`
	ClassificationID  uint            `json:"classification_id"`
	Section           *Section        `json:"section,omitempty"`
	SubSection        *SubSection     `json:"sub_section,omitempty"`
	Classification    *Classification `json:"classification,omitempty"`
	Answers           []*Answer       `json:"answers,omitempty"`
	*QuestionAnalysis `json:"omitempty"`
}

type QuestionAnalysis struct {
	TotalAnswers       int `json:"total_answers,omitempty"`
	TotalIsExist       int `json:"total_is_exist,omitempty"`
	TotalNotExist      int `json:"total_not_exist,omitempty"`
	TotalIsGood        int `json:"total_is_good,omitempty"`
	TotalNotGood       int `json:"total_not_good,omitempty"`
	TotalNa            int `json:"total_na,omitempty"`
	PercentageIsExist  int `json:"percentage_is_exist,omitempty"`
	PercentageNotExist int `json:"percentage_not_exist,omitempty"`
	PercentageIsGood   int `json:"percentage_is_good,omitempty"`
	PercentageNotGood  int `json:"percentage_not_good,omitempty"`
	PercentageNa       int `json:"percentage_na,omitempty"`
}
