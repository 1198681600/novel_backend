package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

func httpWrapper(f func(w http.ResponseWriter, r *http.Request)) func(c *gin.Context) {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func RegisterBookRoutes(router *gin.RouterGroup) {
	router.POST("/create", httpWrapper(BookController.CreateBook))
	router.GET("/get", httpWrapper(BookController.GetBook))
	router.PUT("/update", httpWrapper(BookController.UpdateBook))
	router.DELETE("/delete", httpWrapper(BookController.DeleteBook))
	router.GET("/list", httpWrapper(BookController.ListBooks))
	router.GET("/search", httpWrapper(BookController.SearchBooks))
}

func RegisterNovelRoutes(router *gin.RouterGroup) {
	router.POST("/upsert", httpWrapper(NovelController.UpsertOriginNovel))
	router.GET("/get", httpWrapper(NovelController.GetOriginNovel))
	router.DELETE("/delete", httpWrapper(NovelController.DeleteOriginNovel))
	router.GET("/list", httpWrapper(NovelController.ListOriginNovels))
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
