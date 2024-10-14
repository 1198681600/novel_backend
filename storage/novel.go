package storage

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"novel_backend/global"
	"novel_backend/model"
)

type (
	INovelStorage interface {
		UpsertOriginNovel(bookID int64, chapterID int64, originTitle string, originContent string) error
		GetOriginNovel(bookID int64, chapterID int64) (*model.Novel, error)
		DeleteOriginNovel(bookID int64, chapterID int64) error
		ListOriginNovels(bookID int64, limit, offset int) ([]model.Novel, error)
	}

	novelStorage struct {
	}
)

func newNovelStorage() INovelStorage {
	return &novelStorage{}
}

func (n novelStorage) UpsertOriginNovel(bookID int64, chapterID int64, originTitle string, originContent string) error {
	err := global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "book_id"}, {Name: "chapter_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"chapter_origin_title", "chapter_origin_content"}),
	}).Create(&model.Novel{
		BookID:               bookID,
		ChapterID:            chapterID,
		ChapterOriginTitle:   originTitle,
		ChapterOriginContent: originContent,
	}).Error

	if err != nil {
		global.Logger.Error("UpsertOriginNovel error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID), zap.String("originTitle", originTitle))
	}
	return err
}

func (n novelStorage) GetOriginNovel(bookID int64, chapterID int64) (*model.Novel, error) {
	var novel model.Novel
	err := global.DB.Where("book_id = ? AND chapter_id = ?", bookID, chapterID).First(&novel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		global.Logger.Error("GetOriginNovel error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID))
		return nil, err
	}
	return &novel, nil
}

func (n novelStorage) DeleteOriginNovel(bookID int64, chapterID int64) error {
	err := global.DB.Where("book_id = ? AND chapter_id = ?", bookID, chapterID).Delete(&model.Novel{}).Error
	if err != nil {
		global.Logger.Error("DeleteOriginNovel error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int64("chapterID", chapterID))
	}
	return err
}

func (n novelStorage) ListOriginNovels(bookID int64, limit, offset int) ([]model.Novel, error) {
	var novels []model.Novel
	err := global.DB.Where("book_id = ?", bookID).
		Order("chapter_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&novels).Error

	if err != nil {
		global.Logger.Error("ListOriginNovels error", zap.Error(err), zap.Int64("bookID", bookID), zap.Int("limit", limit), zap.Int("offset", offset))
		return nil, err
	}
	return novels, nil
}
