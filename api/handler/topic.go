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
	err := ctx.MustBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	creatorId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(creatorId))
}
