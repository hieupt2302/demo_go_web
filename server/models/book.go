package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"size:255;not null;index" json:"title"`
	ISBN          string         `gorm:"size:13;unique" json:"isbn"`
	Price         float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	StockQuantity int            `gorm:"default:0" json:"stock_quantity"`
	Publisher     string         `json:"publisher"`
	PublishDate   time.Time      `json:"publish_date"`
	CoverImageURL string         `json:"cover_image_url"`
	Description   string         `gorm:"type:text" json:"description"`
	CategoryID    uint           `json:"category_id"`
	Category      Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	AuthorID      uint           `json:"author_id"`
	Author        Author         `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
