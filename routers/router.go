package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"user-server/controllers"
)

var Db *sql.DB

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	userRouter.GET("/", controllers.GetUser)
	userRouter.POST("/", controllers.CreateUser)

	return r
}
