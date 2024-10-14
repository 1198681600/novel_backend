package model

type Book struct {
	Title    string `gorm:"not null" `
	IsFinish bool   `gorm:"not null" `
	Author   string `gorm:"not null" `
	BaseModel
}
