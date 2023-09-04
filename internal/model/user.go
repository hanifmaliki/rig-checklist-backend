package model

type User struct {
	GormAudit
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	PositionID uint      `json:"position_id"`
	FieldID    uint      `json:"field_id"`
	IsAdmin    bool      `json:"is_admin"`
	IsActive   bool      `json:"is_active"`
	Position   *Position `json:"position,omitempty" gorm:"<-:false"`
	Field      *Field    `json:"field,omitempty" gorm:"<-:false"`
}
