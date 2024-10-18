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

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 token
		token := c.GetHeader("X-Token") // 你可以根据实际情况使用 Authorization 或其他字段

		// 固定的预期 token，可以从配置文件读取
		expectedToken := "345kjcasdhvk120938kmdnbflkjashdf11fds2*((*"

		// 校验 token 是否一致
		if token != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or missing token",
			})
			c.Abort() // 终止请求处理
			return
		}

		// Token 校验通过，继续处理请求
		c.Next()
	}
}

func GetGinHandler() http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery(), CorsHandler())

	// 注册 AuthTokenMiddleware 中间件
	engine.Use(AuthTokenMiddleware())

	RegisterGinHandler(engine)
	return engine.Handler()
}

func RegisterGinHandler(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		novelGroup := api.Group("/novel")
		web.RegisterNovelRoutes(novelGroup)

		bookGroup := api.Group("/book")
		web.RegisterBookRoutes(bookGroup)

		// Add other route groups here as needed
	}
}
