package define

import "novel_backend/model"

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpsertOriginNovelRequest struct {
	BookID               int64  `json:"book_id"`
	ChapterID            int64  `json:"chapter_id"`
	ChapterOriginTitle   string `json:"chapter_origin_title"`
	ChapterOriginContent string `json:"chapter_origin_content"`
}

type UpsertOriginNovelResponse struct {
	BaseResponse
}

type GetOriginNovelRequest struct {
	BookID    int64 `json:"book_id"`
	ChapterID int64 `json:"chapter_id"`
}

type GetOriginNovelResponse struct {
	BaseResponse
	Data *model.Novel `json:"data,omitempty"`
}

type DeleteOriginNovelRequest struct {
	BookID    int64 `json:"book_id"`
	ChapterID int64 `json:"chapter_id"`
}

type DeleteOriginNovelResponse struct {
	BaseResponse
}

type ListOriginNovelsRequest struct {
	BookID   int64 `json:"book_id"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

type ListOriginNovelsResponse struct {
	BaseResponse
	Data []model.Novel `json:"data"`
}
