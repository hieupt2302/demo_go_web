package models

import (
	"time"
)

type User struct {

	ID          int    `gorm:"primaryKey" json:"id"`
	Email       string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`

	PhoneNumber string `json:"phone_number"`
	Role        string `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool   `gorm:"default:false" json:"is_deleted"`
	Cart        []Cart `gorm:"foreignKey:UserID" json:"cart_items,omitempty"`

}
