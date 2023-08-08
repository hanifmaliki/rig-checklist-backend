package model

type CaseStudy struct {
	GormAudit
	Slug           string               `json:"slug" gorm:"uniqueIndex:uidx_case_studies_slug;size:191"`
	CompanyName    string               `json:"company_name"`
	Desc           string               `json:"desc"`
	PersonName     string               `json:"person_name"`
	PersonPosition string               `json:"person_position"`
	Logo           string               `json:"logo"`
	Banner         string               `json:"banner"`
	ImpactDesc     string               `json:"impact_desc"`
	IsActive       bool                 `json:"is_active"`
	ImpactList     []*CaseStudyImpact   `json:"impact_list,omitempty"`
	SolutionList   []*CaseStudySolution `json:"solution_list,omitempty"`
}

type CaseStudyImpact struct {
	GormAudit
	CaseStudyID uint   `json:"case_study_id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Desc        string `json:"desc"`
	Image       string `json:"image"`
	OrderNo     int    `json:"order_no"`
}

type CaseStudySolution struct {
	GormAudit
	CaseStudyID uint     `json:"case_study_id"`
	ProductID   uint     `json:"product_id"`
	OrderNo     int      `json:"order_no"`
	Product     *Product `json:"product,omitempty" gorm:"<-:false"`
}

type CaseStudyMenu struct {
	GormAudit
	CaseStudyID uint       `json:"case_study_id"`
	Type        string     `json:"type"`
	Desc        string     `json:"desc"`
	IsActive    bool       `json:"is_active"`
	OrderNo     int        `json:"order_no"`
	CaseStudy   *CaseStudy `json:"case_study,omitempty" gorm:"<-:false"`
}
