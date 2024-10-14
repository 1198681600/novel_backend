package web

import (
	"encoding/json"
	"net/http"
	"novel_backend/service"
)

var (
	NovelController INovelController
	BookController  IBookController
)

func init() {
	NovelController = newNovelController(service.NovelService)
	BookController = newBookController(service.BookService)
}

// ParseRequest 是一个通用方法，用于解析请求体为指定的结构体
func ParseRequest[T any](req *http.Request, w http.ResponseWriter, dest *T) (err error) {
	err = json.NewDecoder(req.Body).Decode(dest)

	// 解析请求体
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return
}

// SendJSONResponse 是一个通用方法，用于发送 JSON 响应
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	// 发送 JSON 响应
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
