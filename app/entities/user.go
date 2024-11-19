package entities

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id       uuid.UUID `json:"user_id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Email    string    `json:"email" gorm:"not null;unique"`
	Password string    `json:"password" gorm:"not null"`
	Role     string    `json:"role" gorm:"default:'user'"`
}
