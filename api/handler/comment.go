package handler

import (
	"wyatt/api/constant"
	"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var CommentList = func(ctx *gin.Context) {
	var param logic.Comment
	err := ctx.ShouldBindWith(&param, binding.Form)
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

var CommentAdd = func(ctx *gin.Context) {
	var param logic.CommentAdd
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(userId))
}

var CommentDelete = func(ctx *gin.Context) {
	var param logic.CommentDelete
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Delete(userId))
}
