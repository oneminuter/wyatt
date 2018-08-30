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

var UserRegister = func(ctx *gin.Context) {
	//var urParams logic.UserRegister
}

var UserLogin = func(ctx *gin.Context) {
	//可用 账号，手机号，邮箱进行登录

}
