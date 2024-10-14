package storage

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"novel_backend/global"
	"novel_backend/model"
)

type (
	IBookStorage interface {
		CreateBook(title string, isFinish bool, author string) (*model.Book, error)
		GetBook(id int64) (*model.Book, error)
		UpdateBook(id int64, title string, isFinish bool, author string) error
		DeleteBook(id int64) error
		ListBooks(limit, offset int) ([]model.Book, error)
		SearchBooks(query string) ([]model.Book, error)
	}

	bookStorage struct{}
)

func newBookStorage() IBookStorage {
	return &bookStorage{}
}

func (b bookStorage) CreateBook(title string, isFinish bool, author string) (*model.Book, error) {
	book := &model.Book{
		Title:    title,
		IsFinish: isFinish,
		Author:   author,
	}
	err := global.DB.Create(book).Error
	if err != nil {
		global.Logger.Error("CreateBook error", zap.Error(err), zap.String("title", title), zap.String("author", author))
		return nil, err
	}
	return book, nil
}

func (b bookStorage) GetBook(id int64) (*model.Book, error) {
	var book model.Book
	err := global.DB.First(&book, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		global.Logger.Error("GetBook error", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return &book, nil
}

func (b bookStorage) UpdateBook(id int64, title string, isFinish bool, author string) error {
	err := global.DB.Model(&model.Book{}).Where("id = ?", id).Updates(model.Book{
		Title:    title,
		IsFinish: isFinish,
		Author:   author,
	}).Error
	if err != nil {
		global.Logger.Error("UpdateBook error", zap.Error(err), zap.Int64("id", id), zap.String("title", title), zap.String("author", author))
	}
	return err
}

func (b bookStorage) DeleteBook(id int64) error {
	err := global.DB.Delete(&model.Book{}, id).Error
	if err != nil {
		global.Logger.Error("DeleteBook error", zap.Error(err), zap.Int64("id", id))
	}
	return err
}

func (b bookStorage) ListBooks(limit, offset int) ([]model.Book, error) {
	var books []model.Book
	err := global.DB.Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		global.Logger.Error("ListBooks error", zap.Error(err), zap.Int("limit", limit), zap.Int("offset", offset))
		return nil, err
	}
	return books, nil
}

func (b bookStorage) SearchBooks(query string) ([]model.Book, error) {
	var books []model.Book
	err := global.DB.Where("title LIKE ? OR author LIKE ?", "%"+query+"%", "%"+query+"%").Find(&books).Error
	if err != nil {
		global.Logger.Error("SearchBooks error", zap.Error(err), zap.String("query", query))
		return nil, err
	}
	return books, nil
}
