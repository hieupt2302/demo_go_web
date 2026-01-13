package models

type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:255;not null;index:idx_author_name,class:FULLTEXT" json:"name"`
	Bio   string `gorm:"type:text" json:"bio"`
	Books []Book `gorm:"many2many:book_authors;" json:"books,omitempty"`
}