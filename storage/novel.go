package storage

import (
	"go.uber.org/zap"
	"novel_backend/global"
	"novel_backend/model"
)

type (
	INovelStorage interface {
		UpsertOriginNovel(bookID int64, chapterID int64, originTitle string, originContent string) (err error)
	}

	novelStorage struct {
	}
)

func newNovelStorage() INovelStorage {
	return &novelStorage{}
}

func (n novelStorage) UpsertOriginNovel(bookID int64, chapterID int64, originTitle string, originContent string) (err error) {
	err = global.DB.Create(&model.Novel{
		BookID:               bookID,
		ChapterID:            chapterID,
		ChapterOriginTitle:   originTitle,
		ChapterOriginContent: originContent,
	}).Error

	if err != nil {
		global.Logger.Error("UpsertOriginNovel error", zap.Error(err), zap.Any("bookID", bookID), zap.Any("chapterID", chapterID), zap.Any("originTitle", originTitle))
	}
	return
}
