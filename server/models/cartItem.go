package models

import "time"

type CartItem struct {
	ID        int `gorm:"primaryKey" json:"id"`
	BookID    int
	Book      Book
	Quantity  int `gorm:"not null" json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
	isDeleted bool   `gorm:"default:false" json:"is_deleted"`
	Cart []Cart `gorm:"foreignKey:CartItemID" json:"carts,omitempty"`
}
