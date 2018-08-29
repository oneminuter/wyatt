package middleware

import (
	"net/http"
	"strings"
	"wyatt/api/constant"
	//"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"github.com/gin-gonic/gin"
)

var Auth = func(ctx *gin.Context) {
	uuid := ctx.GetHeader("uuid")
	token := ctx.GetHeader("token")

	if "" == strings.TrimSpace(token) && "" == strings.TrimSpace(uuid) {
		//var u logic.User
		//ip := ctx.ClientIP()
		//userId, err := u.AddTempUser(ip)
		//if err != nil {
		//	util.LoggerError(err)
		//	ctx.AbortWithStatusJSON(http.StatusOK, view.SetErr(constant.CreateUserErr))
		//}
		//ctx.Set("userId", userId)
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
