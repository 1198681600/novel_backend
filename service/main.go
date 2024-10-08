package service

import "novel_backend/storage"

var (
	NovelService INovelService
)

func init() {
	NovelService = newNovelService(storage.NovelStorage)
}
