package models

import (
	"time"
)

type User struct {
	ID          int   `gorm:"primaryKey" json:"id"`
	Email       string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password    string `gorm:"not null" json:"-"`
	PhoneNumber int    `json:"phone_number"`
	Role        string `gorm:"type:varchar(50)" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	isDeleted bool `gorm:"default:false" json:"is_deleted"`
	CartItem []CartItem `gorm:"many2many:carts;" json:"cart_items,omitempty"`
}

