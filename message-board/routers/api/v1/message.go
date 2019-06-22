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
	code, err, msg := models.Register(username, password)
	if err != nil{
		c.JSON(code, gin.H{"err" :err, "msg": msg})
	}
	c.JSON(http.StatusOK, gin.H{"err" :err, "msg":"注册成功"})
	logs.Info("user: %s register succeed!", username)
}

func ModifyUser(c *gin.Context)  {
	username := c.MustGet("username").(string)
	newUsername := c.PostForm("username")
	newPassword := c.PostForm("password")
	code, err, msg := models.ModifyUser(username, newUsername, newPassword)
	if err != nil{
		c.JSON(code, gin.H{"err" :err, "msg": msg})
		return
	}
	c.SetCookie("token", "", 0, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg" :"修改账号成功，请重新登录"})
	logs.Info("user: %s is modified!", username)
}

func DeleteUser(c *gin.Context)  {
	username := c.MustGet("username").(string)
	code, err, msg := models.DeleteUser(username)
	if err != nil{
		c.JSON(code, gin.H{"err" :err, "msg": msg})
	}
	c.JSON(http.StatusOK, gin.H{"msg" :"删除账号成功"})
	logs.Info("user: %s is deleted!", username)
}

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code, err, msg := models.Login(username, password)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg": msg})
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"err" : "登录失败","msg":"请稍后再试"})
		logs.Debug(err)
		return
	}
	c.SetCookie("token", token, 3600, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg":"登录成功"})
	logs.Info("user: %s login succeed!", username)
}

func Logout(c *gin.Context)  {
	username := c.MustGet("username").(string)
	c.SetCookie("token", "", 0, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg" :"注销成功"})
	logs.Info("user: %s logout succeed!", username)
}

func GetMessages(c *gin.Context)  {
	page := c.Query("page")
	perPage := c.Query("per_page")
	if (page == "") != (perPage == ""){
		c.JSON(http.StatusBadRequest, gin.H{"err":"请求错误", "msg" : "请求参数不足"})
		return
	}
	messages, code, err, msg := models.GetMessages(page, perPage)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg" : msg})
	}else {
		c.JSON(http.StatusOK, gin.H{"data" : messages})
	}
}

func GetMessageById(c *gin.Context)  {
	messageId := c.Param("message_id")
	message, code, err, msg := models.GetMessageById(messageId)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg" : msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":message})
}

func AddMessage(c *gin.Context)  {
	username := c.MustGet("username").(string)
	title := c.PostForm("title")
	content := c.PostForm("content")
	if title == "" || content == ""{
		c.JSON(http.StatusOK, gin.H{"err" : "添加失败", "msg":"标题和内容都需要填写才能提交"})
		return
	}
	imageUrl, code, err, msg := upload.GetImageUrl(c)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg" : msg})
		return
	}
	message, code, err, msg := models.AddMessage(username, title, content, imageUrl)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg" : msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":message})
}

func ModifyMessage(c *gin.Context) {
	messageId := c.Param("message_id")
	username := c.MustGet("username").(string)
	code, err, msg := models.CheckMessageAuth(username, messageId)
	if err != nil {
		c.JSON(code, gin.H{"err":err, "msg":msg})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	imageUrl, code, err, msg := upload.GetImageUrl(c)
	if err != nil{
		c.JSON(code,gin.H{"err" : err, "msg" : msg})
		return
	}
	message, code, err, msg := models.ModifyMessage(messageId, title, content, imageUrl)
	if err != nil{
		c.JSON(code, gin.H{"err":err, "msg":msg})
	}else {
		c.JSON(http.StatusOK, gin.H{"data":message})
		logs.Info("user %s modify message successfully", c.MustGet("username").(string))
	}
}

func DeleteMessage(c *gin.Context)  {
	messageId := c.Param("message_id")
	username := c.MustGet("username").(string)
	code, err, msg := models.CheckMessageAuth(username, messageId)
	if err != nil {
		c.JSON(code, gin.H{"err":err, "msg":msg})
		return
	}
	code, err, msg = models.DeleteMessage(messageId)
	if err != nil{
		c.JSON(code, gin.H{"err": err, "msg" : msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg" :"删除成功"})
	}

func GetImage(c *gin.Context) {
	name := c.Param("image_name")
	path := setting.AppSetting.ImageSavePath + name
	c.File(path)
}
