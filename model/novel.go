package model

type Novel struct {
	BookID               int64  `gorm:"not null;index;uniqueIndex:idx_book_id_chapter_id" `
	ChapterID            int64  `gorm:"not null;uniqueIndex:idx_book_id_chapter_id" `
	ChapterOriginTitle   string `gorm:"not null" `
	ChapterOriginContent string `gorm:"default:''" `
	ChapterTitle         string `gorm:"not null" `
	ChapterContent       string `gorm:"not null" `
	BaseModel
}
