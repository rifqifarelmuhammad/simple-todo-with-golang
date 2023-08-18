package models

type User struct {
	UID      string `gorm:"primaryKey;column:id" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserID"`
	Timestamp
}
