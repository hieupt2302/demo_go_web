package models

type OrderItem struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	OrderID         uint    `json:"order_id"`
	BookID          uint    `json:"book_id"`
	Quantity        int     `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64 `gorm:"type:decimal(10,2);not null" json:"price_at_purchase"`
	Book            Book    `gorm:"foreignKey:BookID" json:"book"`
}