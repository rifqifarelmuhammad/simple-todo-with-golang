package models

type User struct {
	UID      string `gorm:"primaryKey;column:id;index" json:"id"`
	Email    string `gorm:"unique;index" json:"email"`
	Password string `json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserID"`
	Timestamp
}
