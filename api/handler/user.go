package handler

import (
	"net/http"
	"wyatt/api/logic"

	"strings"
	"wyatt/api/constant"
	"wyatt/api/view"

	"wyatt/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var UserInfo = func(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	var u logic.User
	ctx.JSON(http.StatusOK, u.Info(userId))
}

/*
	注册
	检测 account 和 password 是否为空和是否包含中文字符
*/
var UserRegister = func(ctx *gin.Context) {
	var params logic.UserRegister

	err := ctx.ShouldBindWith(&params, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	if "" == strings.TrimSpace(params.Account) || "" == strings.TrimSpace(params.Password) {
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	if util.IsChineseChar(params.Account) || util.IsChineseChar(params.Password) {
		ctx.JSON(http.StatusOK, view.SetErr(constant.AccountOrPasswordIncludeChinessErr))
	}
	params.Ip = ctx.ClientIP()

	ctx.JSON(http.StatusOK, params.Register())
	return
}

//创建临时用户
var UserTemp = func(ctx *gin.Context) {
	var u logic.User
	ip := ctx.ClientIP()
	ctx.JSON(http.StatusOK, u.AddTempUser(ip))
}

/*
登录
方式可以是：账号，手机号，邮箱
*/
var UserLogin = func(ctx *gin.Context) {
	var params logic.UserLogin

	//可用 账号，手机号，邮箱进行登录
	err := ctx.ShouldBindWith(&params, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	if "" == strings.TrimSpace(params.Password) {
		ctx.JSON(http.StatusOK, view.SetErr(constant.PasswordIsEmptyErr))
		return
	}
	if util.IsChineseChar(params.Account) || util.IsChineseChar(params.Password) {
		ctx.JSON(http.StatusOK, view.SetErr(constant.AccountOrPasswordIncludeChinessErr))
	}

	ctx.JSON(http.StatusOK, params.Login())
	return
}

//修改用户信息
var UserinfoModify = func(ctx *gin.Context) {
	var param logic.UserinfoModify
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Modify(userId))
}

//修改用户账号
var UserAccountModify = func(ctx *gin.Context) {
	var param logic.UserAccountModify
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Modify(userId))
}
