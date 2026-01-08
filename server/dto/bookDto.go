package dto
type BookDto struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	BookName string  `gorm:"type:varchar(200);not null" json:"book_name"`
	Image    string  `gorm:"type:varchar(255);not null" json:"image"`
	Price    float64 `gorm:"not null" json:"price"`
}