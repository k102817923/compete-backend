package routers

import (
	"paracraft-compete-backend/common/setting"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	router.Use(AllowAll())

	// 无需鉴权的路由
	AppNoAuthCheck(router)
	return router
}

// 解决跨域
func AllowAll() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}
