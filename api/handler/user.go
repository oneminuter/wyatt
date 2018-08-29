package handler

import (
	"net/http"
	"wyatt/api/logic"

	"github.com/gin-gonic/gin"
)

var UserInfo = func(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	var u logic.User
	ctx.JSON(http.StatusOK, u.Info(userId))
}
