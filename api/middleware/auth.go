package middleware

import (
	"net/http"
	"strings"
	"wyatt/api/constant"
	"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"github.com/gin-gonic/gin"
)

/*
检测请求头部是否含有
	uuid ：Md5字符串
	token: jwt
如果两个值都为空，则创建一个临时账号
两个值，要么同时存在，要么都不存在
如果通过验证，则将 userId 的值 set 到 ctx 中
*/
var Auth = func(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	if isExitWhite(url) {
		return
	}

	uuid := ctx.GetHeader("uuid")
	token := ctx.GetHeader("token")

	if "" == strings.TrimSpace(token) && "" == strings.TrimSpace(uuid) {
		var u logic.User
		ip := ctx.ClientIP()
		userId, err := u.AddTempUser(ip)
		if err != nil {
			util.LoggerError(err)
			ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.CreateUserErr))
		}
		ctx.Set("userId", userId)
		return
	}

	jwtToken, err := util.ParseToken(token)
	if err != nil && strings.Contains(err.Error(), constant.TokenExpired) {
		ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.AccountExpire))
		return
	}
	if err != nil {
		util.LoggerError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.IllegalAccount))
		return
	}
	if "" == strings.TrimSpace(uuid) || uuid != jwtToken.UUID {
		ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.IllegalRequest))
		return
	}

	ctx.Set("userId", jwtToken.UserId)
}
