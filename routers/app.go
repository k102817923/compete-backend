package routers

import (
	"paracraft-compete-backend/controller"

	"github.com/gin-gonic/gin"
)

func AppNoAuthCheck(router *gin.Engine) {
	app := router.Group("/")
	{
		app.GET("/ping", controller.Test)
		app.GET("/competes", controller.GetCompeteList)
	}

}
