package middleware

import (
	"net/http"
	"wyatt/api/constant"
	"wyatt/api/view"

	"github.com/gin-gonic/gin"
)

var whiteUrlList = []string{
	"/ping",
}
var MustLogin = func(ctx *gin.Context) {
	url := ctx.Request.URL.Path

	var isExit bool
	for _, v := range whiteUrlList {
		if v == url {
			isExit = true
			break
		}
	}

	if isExit {
		return
	}

	if 0 == ctx.GetInt64("userId") {
		ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.IllegalAccount))
	}
}
