package models

import (
	"time"
)

type Order struct {
	ID              uint        `gorm:"primaryKey" json:"id"`
	UserID          uint        `json:"user_id"`
	OrderDate       time.Time   `gorm:"autoCreateTime" json:"order_date"`
	TotalAmount     float64     `gorm:"type:decimal(10,2)" json:"total_amount"`
	Status          string      `gorm:"type:enum('pending','shipped','delivered','cancelled');default:'pending'" json:"status"`
	PaymentMethod   string      `gorm:"type:varchar(50)" json:"payment_method"`
    PaymentStatus   string      `gorm:"type:varchar(50);default:'unpaid'" json:"payment_status"` 
    TxnID           string      `gorm:"type:varchar(100)" json:"txn_id"`
	ShippingAddress string      `gorm:"type:text" json:"shipping_address"`
	OrderItems      []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}