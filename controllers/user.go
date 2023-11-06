package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-server/models"
	"user-server/services"
)

func GetUser(c *gin.Context) {
	userList := services.GetUser()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": userList,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	// 将 req 的 body 数据按照 JSON 的格式解析到 User 结构体变量中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	services.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "",
		"data": user,
	})
}
