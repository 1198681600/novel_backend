package service

import (
	"go.uber.org/zap"
	"novel_backend/define"
	"novel_backend/global"
	"novel_backend/model"
	"novel_backend/storage"
)

type (
	INovelService interface {
		UpsertOriginNovel(req define.UpsertOriginNovelRequest) error
		GetOriginNovel(bookID, chapterID int64) (*model.Novel, error)
		DeleteOriginNovel(bookID, chapterID int64) error
		ListOriginNovels(bookID int64) ([]string, error)
	}

	novelService struct {
		novelStorage storage.INovelStorage
	}
)

func newNovelService(novelStorage storage.INovelStorage) INovelService {
	return &novelService{
		novelStorage: novelStorage,
	}
}

func (n novelService) UpsertOriginNovel(req define.UpsertOriginNovelRequest) error {
	err := n.novelStorage.UpsertOriginNovel(req.BookID, req.ChapterID, req.ChapterOriginTitle, req.ChapterOriginContent)
	if err != nil {
		global.Logger.Error("UpsertOriginNovel error", zap.Error(err), zap.Any("request", req))
		return err
	}
	global.Logger.Info("Successfully upserted origin novel", zap.Int64("bookID", req.BookID), zap.Int64("chapterID", req.ChapterID))
	return nil
}

func (n novelService) GetOriginNovel(bookID, chapterID int64) (*model.Novel, error) {
	novel, err := n.novelStorage.GetOriginNovel(bookID, chapterID)
	if err != nil {
		global.Logger.Error("GetOriginNovel error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID))
		return nil, err
	}
	return novel, nil
}

func (n novelService) DeleteOriginNovel(bookID, chapterID int64) error {
	err := n.novelStorage.DeleteOriginNovel(bookID, chapterID)
	if err != nil {
		global.Logger.Error("DeleteOriginNovel error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID))
		return err
	}
	global.Logger.Info("Successfully deleted origin novel", zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID))
	return nil
}

func (n novelService) ListOriginNovels(bookID int64) ([]string, error) {
	novels, err := n.novelStorage.ListOriginNovels(bookID)
	if err != nil {
		global.Logger.Error("ListOriginNovels error", zap.Error(err), zap.Int64("bookID", bookID))
		return nil, err
	}
	return novels, nil
}
