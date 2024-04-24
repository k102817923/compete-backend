package controller

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	HttpOkResponse(c, map[string]string{"message": "pong!"})
}
