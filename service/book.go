package service

import (
	"go.uber.org/zap"
	"novel_backend/global"
	"novel_backend/model"
	"novel_backend/storage"
)

type (
	IBookService interface {
		CreateBook(title string, isFinish bool, author string, image string, introduction string, category model.Category, metadata map[string]string) (*model.Book, error)
		GetBook(id int64) (*model.Book, error)
		UpdateBook(id int64, title string, isFinish bool, author string) error
		DeleteBook(id int64) error
		ListBooks(page, pageSize int) ([]model.Book, error)
		SearchBooks(query string) ([]model.Book, error)
	}

	bookService struct {
		bookStorage storage.IBookStorage
	}
)

func newBookService(bookStorage storage.IBookStorage) IBookService {
	return &bookService{
		bookStorage: bookStorage,
	}
}

func (b bookService) CreateBook(title string, isFinish bool, author string, image string, introduction string, category model.Category, metadata map[string]string) (*model.Book, error) {
	book, err := b.bookStorage.CreateBook(title, isFinish, author, image, introduction, category, metadata)
	if err != nil {
		global.Logger.Error("CreateBook error", zap.Error(err))
		return nil, err
	}
	global.Logger.Info("Successfully created book", zap.Int64("bookID", book.ID), zap.String("title", book.Title))
	return book, nil
}

func (b bookService) GetBook(id int64) (*model.Book, error) {
	book, err := b.bookStorage.GetBook(id)
	if err != nil {
		global.Logger.Error("GetBook error", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return book, nil
}

func (b bookService) UpdateBook(id int64, title string, isFinish bool, author string) error {
	err := b.bookStorage.UpdateBook(id, title, isFinish, author)
	if err != nil {
		global.Logger.Error("UpdateBook error", zap.Error(err))
		return err
	}
	global.Logger.Info("Successfully updated book")
	return nil
}

func (b bookService) DeleteBook(id int64) error {
	err := b.bookStorage.DeleteBook(id)
	if err != nil {
		global.Logger.Error("DeleteBook error", zap.Error(err), zap.Int64("id", id))
		return err
	}
	global.Logger.Info("Successfully deleted book", zap.Int64("bookID", id))
	return nil
}

func (b bookService) ListBooks(page, pageSize int) ([]model.Book, error) {
	offset := (page - 1) * pageSize
	books, err := b.bookStorage.ListBooks(pageSize, offset)
	if err != nil {
		global.Logger.Error("ListBooks error", zap.Error(err), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return nil, err
	}
	return books, nil
}

func (b bookService) SearchBooks(query string) ([]model.Book, error) {
	books, err := b.bookStorage.SearchBooks(query)
	if err != nil {
		global.Logger.Error("SearchBooks error", zap.Error(err), zap.String("query", query))
		return nil, err
	}
	return books, nil
}
