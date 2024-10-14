package service

import "novel_backend/storage"

var (
	NovelService INovelService
	BookService  IBookService
)

func init() {
	NovelService = newNovelService(storage.NovelStorage)
	BookService = newBookService(storage.BookStorage)
}
