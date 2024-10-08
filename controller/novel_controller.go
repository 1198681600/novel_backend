package web

import (
	"net/http"
	"novel_backend/define"
	"novel_backend/service"
)

type (
	INovelController interface {
		UpsertOriginNovel(w http.ResponseWriter, r *http.Request)
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

func (c novelController) UpsertOriginNovel(w http.ResponseWriter, req *http.Request) {
	// 创建一个请求体的实例
	var request define.UpsertOriginNovelRequest

	// 解析请求体
	if err := ParseRequest(req, w, &request); err != nil {
		return
	}

	// 调用服务层的方法
	if err := c.NovelService.UpsertOriginNovel(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 创建一个响应体的实例
	response := define.UpsertOriginNovelResponse{}

	// 发送 JSON 响应
	SendJSONResponse(w, http.StatusOK, response)
}
