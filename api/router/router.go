package router

import (
	"wyatt/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("ping", handler.Ping)
	//用户
	user := server.Group("/user")
	user.GET("/info", handler.UserInfo)
	user.POST("/register", handler.UserRegister)
	user.POST("/login", handler.UserLogin)

}
