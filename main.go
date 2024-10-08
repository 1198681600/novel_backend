package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"net/http"
	web "novel_backend/controller"
	"novel_backend/global"
	"novel_backend/model"
)

func main() {

	model.RegisterTables()

	g := &run.Group{}

	// Add Http Server
	httpServer := global.CreateHttpServer()
	g.Add(HttpServerExecute(httpServer), func(error) {
		_ = httpServer.Close()
	})

	err := g.Run()
	if err != nil {
		global.Logger.Error("run server error", zap.Error(err))
		return
	}
}

func HttpServerExecute(httpServer *http.Server) func() error {
	return func() error {
		httpServer.Handler = GetGinHandler()
		httpServer.Addr = ":7899"
		global.Logger.Info("starting HTTP server", zap.String("addr", httpServer.Addr))
		return httpServer.ListenAndServe()
	}
}

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header(
			"Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
		) // Cross-domain key settings allow browsers to resolve.
		c.Header(
			"Access-Control-Max-Age",
			"172800",
		) // Cache request information in seconds.
		c.Header(
			"Access-Control-Allow-Credentials",
			"false",
		) //  Whether cross-domain requests need to carry cookie information, the default setting is true.
		c.Header(
			"content-type",
			"application/json",
		) // Set the return format to json.
		// Release all option pre-requests
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, "Options Request!")
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetGinHandler() http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery(), CorsHandler())
	RegisterGinHandler(engine)
	return engine.Handler()
}

func httpWrapper(f func(w http.ResponseWriter, r *http.Request)) func(c *gin.Context) {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func RegisterGinHandler(engine *gin.Engine) {
	engine.POST("/novel/upsert_origin_novel", httpWrapper(web.NovelController.UpsertOriginNovel))
}
