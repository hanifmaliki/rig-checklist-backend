package model

import "time"

type Form struct {
	GormAudit
	FormNo       string    `json:"form_no"`
	RigID        uint      `json:"rig_id"`
	LocationID   uint      `json:"location_id"`
	WellID       uint      `json:"well_id"`
	FieldID      uint      `json:"field_id"`
	ActivityID   uint      `json:"activity_id"`
	Date         time.Time `json:"date"`
	IsInspected1 bool      `json:"is_inspected_1"`
	IsInspected2 bool      `json:"is_inspected_2"`
	IsInspected3 bool      `json:"is_inspected_3"`
	IsInspected4 bool      `json:"is_inspected_4"`
	IsChecked1   bool      `json:"is_checked_1"`
	IsChecked2   bool      `json:"is_checked_2"`
	IsChecked3   bool      `json:"is_checked_3"`
	IsApproved   bool      `json:"is_approved"`
	Operable     *bool     `json:"operable"`
	Rig          *Rig      `json:"rig,omitempty" gorm:"<-:false"`
	Location     *Location `json:"location,omitempty" gorm:"<-:false"`
	Well         *Well     `json:"well,omitempty" gorm:"<-:false"`
	Field        *Field    `json:"field,omitempty" gorm:"<-:false"`
	Activity     *Activity `json:"activity,omitempty" gorm:"<-:false"`
	Answers      []*Answer `json:"answers,omitempty"`
}
