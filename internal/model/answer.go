package model

type Answer struct {
	GormAudit
	FormID     uint      `json:"form_id"`
	QuestionID uint      `json:"question_id"`
	IsExist    *bool     `json:"is_exist"`
	IsGood     *bool     `json:"is_good"`
	NA         bool      `json:"na"`
	Note       string    `json:"note"`
	Question   *Question `json:"question,omitempty" gorm:"<-:false"`
	Photos     *[]Photo  `json:"photos,omitempty" gorm:"<-:false"`
}
