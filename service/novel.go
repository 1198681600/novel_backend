package service

import (
	"go.uber.org/zap"
	"novel_backend/define"
	"novel_backend/global"
	"novel_backend/storage"
)

type (
	INovelService interface {
		UpsertOriginNovel(req define.UpsertOriginNovelRequest) (err error)
	}

	novelService struct {
		novelStorage storage.INovelStorage
	}
)

func newNovelService(novelStorage storage.INovelStorage) INovelService {
	s := &novelService{
		novelStorage: novelStorage,
	}
	return s
}

func (n novelService) UpsertOriginNovel(req define.UpsertOriginNovelRequest) (err error) {
	err = n.novelStorage.UpsertOriginNovel(req.BookID, req.ChapterID, req.ChapterOriginTitle, req.ChapterOriginContent)
	if err != nil {
		return err
	}
	global.Logger.Info("成功写入原文", zap.Any("bookID", req.BookID), zap.Any("chapterID", req.ChapterID), zap.Any("originTitle", req.ChapterOriginTitle))
	return
}
