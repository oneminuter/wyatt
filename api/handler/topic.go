package handler

import (
	"net/http"
	"wyatt/api/constant"
	"wyatt/api/logic"
	"wyatt/api/view"
	"wyatt/util"

	"strings"

	"errors"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/json-iterator/go"
)

var TopicList = func(ctx *gin.Context) {
	var (
		param logic.Topic
	)

	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	bytes, _ := jsoniter.Marshal(param)
	log.Println(string(bytes))
	if "" == param.CId {
		util.LoggerError(errors.New("社区流水号为空"))
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
	if "" != strings.TrimSpace(tId) {
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	ctx.JSON(http.StatusOK, t.Detail(tId))
	return
}

var TopicCollectAdd = func(ctx *gin.Context) {
	var param logic.TopicCollect
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(userId))
}

var TopicCollectCancel = func(ctx *gin.Context) {
	var param logic.TopicCollect
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Cancel(userId))
}

var TopicCollectList = func(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var param logic.TopicCollect
	ctx.JSON(http.StatusOK, param.List(userId))
}
