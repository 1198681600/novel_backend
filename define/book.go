package define

import "novel_backend/model"

type CreateBookRequest struct {
	Title        string            `json:"title"`        // 书名
	IsFinish     bool              `json:"isFinish"`     // 是否完结
	Author       string            `json:"author"`       // 作者
	Image        string            `json:"image"`        // 封面图片URL
	Introduction string            `json:"introduction"` // 简介
	Category     model.Category    `json:"category"`     // 分类
	Metadata     map[string]string `json:"metadata"`     // 其他元数据
}

type CreateBookResponse struct {
	BaseResponse
	Data *model.Book `json:"data,omitempty"`
}

type GetBookRequest struct {
	ID int64 `json:"id"`
}

type GetBookResponse struct {
	BaseResponse
	Data *model.Book `json:"data,omitempty"`
}

type UpdateBookRequest struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	IsFinish bool   `json:"isFinish"`
	Author   string `json:"author"`
}

type UpdateBookResponse struct {
	BaseResponse
}

type DeleteBookRequest struct {
	ID int64 `json:"id"`
}

type DeleteBookResponse struct {
	BaseResponse
}

type ListBooksRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ListBooksResponse struct {
	BaseResponse
	Data []model.Book `json:"data"`
}

type SearchBooksRequest struct {
	Query string `json:"query"`
}

type SearchBooksResponse struct {
	BaseResponse
	Data []model.Book `json:"data"`
}
