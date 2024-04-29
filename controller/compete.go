package controller

import (
	"paracraft-compete-backend/common/models"

	"github.com/gin-gonic/gin"
)

func GetCompeteList(c *gin.Context) {
	data := make(map[string]interface{})
	data["raw"] = models.GetCompeteList()
	HttpOkResponse(c, data)
}
