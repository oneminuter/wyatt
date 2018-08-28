package router

import (
	"wyatt/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	user := server.Group("/user")
	user.GET("/info", handler.UserInfo)

}
