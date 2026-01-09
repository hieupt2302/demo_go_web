package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Fullname     string    `gorm:"size:255" json:"fullname"`
	Email        string    `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Phone        string    `gorm:"size:15" json:"phone"`
	Address      string    `gorm:"type:text" json:"address"`
	Role         string    `gorm:"type:enum('customer', 'admin');default:'customer'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
