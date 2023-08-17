package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	IsDeleted   bool      `gorm:"default:false" json:"is_deleted"`
	UserID      uuid.UUID
}
