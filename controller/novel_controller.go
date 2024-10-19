package web

import (
	"net/http"

	"novel_backend/define"
	"novel_backend/service"
)

type (
	INovelController interface {
		UpsertOriginNovel(w http.ResponseWriter, r *http.Request)
		GetOriginNovel(w http.ResponseWriter, r *http.Request)
		DeleteOriginNovel(w http.ResponseWriter, r *http.Request)
		ListOriginNovels(w http.ResponseWriter, r *http.Request)
	}

	novelController struct {
		NovelService service.INovelService
	}
)

func newNovelController(NovelService service.INovelService) INovelController {
	return &novelController{
		NovelService: NovelService,
	}
}

func (c novelController) UpsertOriginNovel(w http.ResponseWriter, r *http.Request) {
	var req define.UpsertOriginNovelRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	err := c.NovelService.UpsertOriginNovel(req)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.UpsertOriginNovelResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.UpsertOriginNovelResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
	})
}

func (c novelController) GetOriginNovel(w http.ResponseWriter, r *http.Request) {
	var req define.GetOriginNovelRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	novel, err := c.NovelService.GetOriginNovel(req.BookID, req.ChapterID)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.GetOriginNovelResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.GetOriginNovelResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
		Data:         novel,
	})
}

func (c novelController) DeleteOriginNovel(w http.ResponseWriter, r *http.Request) {
	var req define.DeleteOriginNovelRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	err := c.NovelService.DeleteOriginNovel(req.BookID, req.ChapterID)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.DeleteOriginNovelResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.DeleteOriginNovelResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
	})
}

func (c novelController) ListOriginNovels(w http.ResponseWriter, r *http.Request) {
	var req define.ListOriginNovelsRequest
	if err := ParseRequest(r, w, &req); err != nil {
		return
	}

	novels, err := c.NovelService.ListOriginNovels(req.BookID)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, define.ListOriginNovelsResponse{
			BaseResponse: define.BaseResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, define.ListOriginNovelsResponse{
		BaseResponse: define.BaseResponse{Code: http.StatusOK, Message: "Success"},
		Data:         novels,
	})
}
