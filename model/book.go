package model

type Book struct {
	Title    string   `gorm:"not null" `
	IsFinish bool     `gorm:"not null" `
	Author   string   `gorm:"not null" `
	Category Category `gorm:"not null" `
	BaseModel
}

type Category int64

const (
	BookCategoryXuanHuan Category = iota
)
