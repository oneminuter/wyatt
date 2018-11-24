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
	var param logic.StoryRole
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(userId))
}
var RoleModify = func(ctx *gin.Context) {}
var RoleDelete = func(ctx *gin.Context) {}
var RoleList = func(ctx *gin.Context) {}
var RoleInfo = func(ctx *gin.Context) {}
