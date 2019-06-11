package v1

import (
"fmt"
"github.com/gin-gonic/gin"
"message-board/models"
"net/http"
)

func Register(c *gin.Context)  {
	userId := c.Request.FormValue("user_id")
	userPass := c.Request.FormValue("user_pass")
	fmt.Println(userId, userPass)
	err := models.Register(userId, userPass)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err})
	}
}

func Login(c *gin.Context)  {
	userId := c.Request.FormValue("user_id")
	userPass := c.Request.FormValue("user_pass")
	err := models.Login(userId, userPass)
	if err == nil{
		c.SetCookie("user_id", userId, 3600, "/", "127.0.0.1", false, true)
		c.SetCookie("user_pass", userPass, 3600, "/", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, gin.H{"err" :err})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}
}

func GetAllMessages(c *gin.Context)  {
	ms, err := models.GetAllMessages()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"messages" : ms, "err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"messages" : ms, "err" :err})
	}
}

func AddMessage(c *gin.Context)  {
	userId, err := c.Cookie("user_id")
	userPass, err := c.Cookie("user_pass")
	mCont := c.Request.FormValue("m_cont")
	mId, err := models.AddMessage(userId, userPass, mCont)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"m_id":mId, "err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"m_id":mId, "err" :err})
	}
}

func DeleteMessage(c *gin.Context)  {
	userId, err := c.Cookie("user_id")
	userPass, err := c.Cookie("user_pass")
	mId := c.Query("m_id")
	fmt.Println(mId)
	err = models.DeleteMessage(userId, userPass, mId)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"err" :err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"err" :err})
	}
}

