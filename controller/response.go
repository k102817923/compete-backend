package controller

import (
	"net/http"
	e "paracraft-compete-backend/common/e"
	"time"

	"github.com/gin-gonic/gin"
)

func HttpOkResponse(c *gin.Context, data interface{}) {
	msg := e.GetMsg(e.SUCCESS)
	if data == nil {
		data = []int{}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  msg,
		"data": data,
		"time": int(time.Now().Unix()),
	})
}
