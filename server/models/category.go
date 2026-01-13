package models

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null;index:idx_category_name,class:FULLTEXT" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Books       []Book `gorm:"many2many:book_categories;" json:"books,omitempty"`
}