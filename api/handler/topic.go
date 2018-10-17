package handler

import (
	"net/http"
	"strconv"
	"wyatt/api/constant"
	"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var TopicList = func(ctx *gin.Context) {
	var (
		param logic.Topic
		err   error
	)

	cId := ctx.Query("cId")
	if "" == cId {
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	param.CId, err = strconv.ParseInt(cId, 10, 64)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	ctx.JSON(http.StatusOK, param.List())
}

var TopicAdd = func(ctx *gin.Context) {
	var param logic.TopicAdd
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	creatorId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(creatorId))
}

var TopicDelete = func(ctx *gin.Context) {
	var param logic.TopicDelete
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Delete(userId))
}

var TopicModify = func(ctx *gin.Context) {
	var param logic.TopicModify
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Modify(userId))
}

var TopicDetail = func(ctx *gin.Context) {
	var t logic.Topic
	tId := ctx.Query("tId")
	topicId, err := strconv.ParseInt(tId, 10, 64)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	ctx.JSON(http.StatusOK, t.Detail(topicId))
	return
}
