package v1

import (
"github.com/astaxie/beego/logs"
"github.com/gin-gonic/gin"
"message-board/models"
"message-board/pkg/setting"
"message-board/pkg/upload"
"message-board/pkg/util"
"net/http"
)

func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := models.Register(username, password)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(),"msg":"注册失败"})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err, "msg":"注册成功"})
		logs.Info("user: %s register succeed!", username)
	}
}

func DeleteUser(c *gin.Context)  {
	username := c.MustGet("username").(string)
	err := models.DeleteUser(username)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"msg" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"msg" :"账号被删除成功"})
		logs.Info("user: %s is deleted!", username)
	}
}

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := models.Login(username, password)
	if err == nil{
		token, err := util.GenerateToken(username, password)
		if err == nil{
			c.SetCookie("token", token, 3600, "/", "127.0.0.1", false, true)
			c.JSON(http.StatusOK, gin.H{"err" :err, "msg":"登录成功"})
			logs.Info("user: %s login succeed!", username)
		}else {
			c.JSON(http.StatusOK, gin.H{"err" :err.Error(),"msg":"登录失败"})
			logs.Debug(err)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(),"msg":"登录失败"})
		logs.Info("user: %s login failed!", username)
	}
}

func Logout(c *gin.Context)  {
	username := c.MustGet("username").(string)
	logs.Info("user: %s logout succeed!", username)
	c.SetCookie("token", "", 0, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg" :"注销成功"})
}

func GetMessages(c *gin.Context)  {
	page := c.Query("page")
	perPage := c.Query("per_page")
	if (page == "") != (perPage == ""){
		c.JSON(http.StatusOK, gin.H{"msg" : "请求错误，请求参数不足"})
		return
	}
	messages, err := models.GetMessages(page, perPage)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "data" : messages})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err, "data" : messages})
	}
}

func GetMessageById(c *gin.Context)  {
	messageId := c.Param("message_id")
	message, err := models.GetMessageById(messageId)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(),"msg":"获取失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err" :err,"data":message})
}

func AddMessage(c *gin.Context)  {
	username := c.MustGet("username").(string)
	title := c.PostForm("title")
	content := c.PostForm("content")
	imageUrl, err := upload.GetImageUrl(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "msg":"添加失败"})
		return
	}
	message, err := models.AddMessage(username, title, content, imageUrl)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "msg":"添加失败"})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err, "msg":"添加成功", "data":message})
	}
}

func ModifyMessage(c *gin.Context) {
	messageId := c.Param("message_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	imageUrl, err := upload.GetImageUrl(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "msg":"修改失败"})
		return
	}
	message, err := models.ModifyMessage(messageId, title, content, imageUrl)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "msg": "修改失败"})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err,"msg":"修改成功","data":message})
	}
}

func DeleteMessage(c *gin.Context)  {
	messageId := c.Param("id")
	username := c.MustGet("username").(string)
	err := models.DeleteMessage(messageId, username)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error(), "msg":"删除失败"})
	}else {
		c.JSON(http.StatusOK, gin.H{"msg" :"删除成功"})
	}
}

func GetImage(c *gin.Context) {
	name := c.Param("name")
	path := setting.AppSetting.ImageSavePath + name
	c.File(path)
}
