package models

import "time"

type Cart struct {
	CartItemID uint `gorm:"primaryKey" json:"cart_item_id"`
	UserID     uint `gorm:"primaryKey" json:"user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
