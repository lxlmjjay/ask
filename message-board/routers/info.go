package routers

import (
"github.com/gin-gonic/gin"
"net/http"
)

func Info(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"links":[]gin.H{
		{ "rel":"获取指定名称的图片", "method": "get", "href":"/api/v1/images/xxx" } ,

		{ "rel":"注册", "method": "post", "href":"/api/v1/user" } ,
		{ "rel":"修改账号" , "method": "put（携带有效token）", "href":"/api/v1/user" },
		{ "rel":"删号", "method": "delete（携带有效token）", "href":"/api/v1/user" },

		{ "rel":"登录", "method": "delete", "href":"/api/v1/auth" },
		{ "rel":"注销", "method": "delete", "href":"/api/v1/auth" },

		{ "rel":"获取message列表", "method": "get", "href":"/api/v1/messages?page=xxx&per_page=xxx" },
		{ "rel":"上传一个message", "method": "post", "href":"/api/v1/messages" },
		{ "rel":"获取指定ID的message", "method": "get", "href":"/api/v1/messages/xxx" },
		{ "rel":"修改指定ID的message", "method": "put", "href":"/api/v1/messages/xxx" },
		{ "rel":"删除指定ID的message", "method": "delete", "href":"/api/v1/messages/xxx" },
	}})
}
