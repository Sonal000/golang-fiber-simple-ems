package entities

import "github.com/gofrs/uuid"

type Employee struct {
	Id         uuid.UUID `json:"emp_id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name       string    `json:"name" gorm:"not null"`
	Position   string    `json:"position" gorm:"not null"`
	Department string    `json:"department" gorm:"not null"`
	Salary     float64   `json:"salary" gorm:"not null"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null;unique"`
	User       User      `json:"user" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
}
