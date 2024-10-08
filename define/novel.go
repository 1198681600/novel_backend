package define

type UpsertOriginNovelRequest struct {
	BookID               int64  `json:"book_id"`
	ChapterID            int64  `json:"chapter_id"`
	ChapterOriginTitle   string `json:"chapter_origin_title"`
	ChapterOriginContent string `json:"chapter_origin_content"`
}

type UpsertOriginNovelResponse struct {
}
