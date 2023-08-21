package models

type Todo struct {
	ID          string `gorm:"primaryKey;index" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `gorm:"default:false" json:"is_completed"`
	Timestamp
	IsDeleted bool   `gorm:"default:false" json:"is_deleted"`
	UserID    string `gorm:"index"`
}
