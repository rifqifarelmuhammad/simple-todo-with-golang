package models

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `gorm:"primary key" json:"uid" db:"uid"`
	Email    string    `gorm:"unique" json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
}
