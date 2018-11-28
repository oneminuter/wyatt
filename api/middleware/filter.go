package middleware

import (
	"net/http"
	"wyatt/api/constant"
	"wyatt/api/view"

	"github.com/gin-gonic/gin"
)

//必须登录检测
var MustLogin = func(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	if isExitWhite(url) {
		return
	}

	if 0 == ctx.GetInt64("userId") {
		ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.IllegalAccount))
		return
	}
}
