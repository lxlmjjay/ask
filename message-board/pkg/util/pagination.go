package util

import (
	"github.com/gin-gonic/gin"
	"message-board/pkg/setting"
	"strconv"
)

func GetPage(c *gin.Context) (result int) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
