package routes

import (
	"go-webapp/config"
	"go-webapp/handle"
	"go-webapp/middleware/cors"
	"go-webapp/middleware/log"
	"go-webapp/middleware/request"

	"github.com/gin-gonic/gin"
	// proxy "github.com/chenhg5/gin-reverseproxy"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//InitRouter Initialise router
func InitRouter() *gin.Engine {
	route := gin.New()

	//route.Use(gzip.Gzip(gzip.DefaultCompression))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if config.GetEnv().DEBUG {
		route.Use(gin.Logger()) // Used in development mode, console print request records
	}

	route.Use(log.JSONLogMiddleware())
	route.Use(gin.Recovery())
	route.Use(request.RequestID(request.RequestIDOptions{AllowSetting: false}))

	route.Use(cors.CORS(cors.CORSOptions{}))
	route.Use(handle.Errors()) // Error handling

	registerAPIRouter(route)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return route
}
