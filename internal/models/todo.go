package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	Timestamp
	IsDeleted bool `gorm:"default:false" json:"is_deleted"`
	UserID    uuid.UUID
}
