package models

import "time"

type Cart struct {
	CartItemID int      `gorm:"primaryKey" json:"cart_item_id"`
	UserID     int      `gorm:"primaryKey" json:"user_id"`
	User       User     `gorm:"foreignKey:UserID;references:ID"`
	CartItem   CartItem `gorm:"foreignKey:CartItemID;references:ID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
