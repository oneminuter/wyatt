package handler

import (
	"net/http"
	"wyatt/api/constant"
	"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var RoleAdd = func(ctx *gin.Context) {
	var param logic.StoryRoleAdd
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(userId))
}
var RoleModify = func(ctx *gin.Context) {
	var param logic.StoryRoleModify
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Modify(userId))
}
var RoleDelete = func(ctx *gin.Context) {
	var param logic.StoryRole
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Delete(userId))
}
var RoleList = func(ctx *gin.Context) {
	var param logic.StoryRoleList
	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	//判断参数是否合法
	if 0 > param.Page || 0 > param.Limit {
		ctx.JSON(http.StatusOK, view.SetErr(constant.QueryPageOrLimit))
		return
	}
	if constant.MAX_QUERY_COUNT < param.Limit {
		param.Limit = constant.MAX_QUERY_COUNT
	}
	ctx.JSON(http.StatusOK, param.List())
}
var RoleInfo = func(ctx *gin.Context) {
	var param logic.StoryRole
	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	ctx.JSON(http.StatusOK, param.Info())
}
