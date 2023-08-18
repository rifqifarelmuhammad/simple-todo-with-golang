package models

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `gorm:"primaryKey;column:id" json:"id"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password"`
	Todos    []Todo    `gorm:"foreignKey:UserID"`
	Timestamp
}
