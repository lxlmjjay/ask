package v1

import (
	"github.com/gin-gonic/gin"
	"message-board/models"
	"message-board/pkg/util"
	"net/http"
	"strconv"
)

func Register(c *gin.Context)  {
	userId := c.PostForm("user_id")
	userPass := c.PostForm("user_pass")
	err := models.Register(userId, userPass)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err})
	}
}

func Login(c *gin.Context)  {
	userIdStr := c.PostForm("user_id")
	userPassStr := c.PostForm("user_pass")
	err := models.Login(userIdStr, userPassStr)
	if err == nil{
		userId,_ := strconv.Atoi(userIdStr)
		userPass,_ := strconv.Atoi(userPassStr)
		token, err := util.GenerateToken(userId, userPass)
		if err == nil{
			c.SetCookie("token", token, 3600, "/", "127.0.0.1", false, true)
			c.JSON(http.StatusOK, gin.H{"err" :err})
		}else {
			c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}
}

func GetAllMessages(c *gin.Context)  {
	ms, err := models.GetAllMessages()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"user_id":c.MustGet("userId").(int),"messages" : ms, "err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"user_id":c.MustGet("userId").(int),"messages" : ms, "err" :err})
	}
}

func AddMessage(c *gin.Context)  {
	userId := c.MustGet("userId").(int)
	mCont := c.PostForm("m_cont")
	mId, err := models.AddMessage(userId, mCont)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"user_id":c.MustGet("userId").(int),"m_id":mId, "err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"user_id":c.MustGet("userId").(int),"m_id":mId, "err" :err})
	}
}

//func EditMessage(c *gin.Context) {
//	userId, err := c.Cookie("user_id")
//	userPass, err := c.Cookie("user_pass")
//	mId := c.Request.FormValue("m_id")
//	mHead := c.Request.FormValue("m_head")
//	mCont := c.Request.FormValue("m_cont")
//	err = models.EditMessage(userId, userPass, mId, mHead mCont)
//	if err != nil{
//		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
//	}else {
//		c.JSON(http.StatusOK, gin.H{"err" :err})
//	}
//}

func DeleteMessage(c *gin.Context)  {
	mId := c.Query("m_id")
	err := models.DeleteMessage(mId)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err})
	}
}

