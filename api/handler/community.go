package handler

import (
	"net/http"
	"wyatt/api/logic"

	"wyatt/api/constant"
	"wyatt/api/view"
	"wyatt/util"

	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//全部社区列表
var CommunityListAll = func(ctx *gin.Context) {
	var lc logic.Community
	err := ctx.ShouldBindQuery(&lc)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	//判断参数是否合法
	if 0 > lc.Page || 0 > lc.Limit {
		ctx.JSON(http.StatusOK, view.SetErr(constant.QueryPageOrLimit))
		return
	}
	if constant.MAX_QUERY_COUNT < lc.Limit {
		lc.Limit = constant.MAX_QUERY_COUNT
	}

	ctx.JSON(http.StatusOK, lc.ListAll())
}

//我加入的社区列表
var CommunityListMy = func(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var (
		ljc logic.JoinedCommunity
		err error
	)
	page := ctx.DefaultQuery("page", "0")
	ljc.Page, err = strconv.Atoi(page)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	maxQueryCount := strconv.Itoa(constant.MAX_QUERY_COUNT)
	limit := ctx.DefaultQuery("limit", maxQueryCount)
	ljc.Limit, err = strconv.Atoi(limit)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	//判断参数是否合法
	if 0 > ljc.Page || 0 > ljc.Limit {
		ctx.JSON(http.StatusOK, view.SetErr(constant.QueryPageOrLimit))
		return
	}
	if constant.MAX_QUERY_COUNT < ljc.Limit {
		ljc.Limit = constant.MAX_QUERY_COUNT
	}

	ctx.JSON(http.StatusOK, ljc.MyList(userId))
}

//加入社区
var Join = func(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	//获取社区号，创建社区的的时间戳(秒)
	var jc logic.JoinedCommunity
	err := ctx.ShouldBindWith(&jc, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	ctx.JSON(http.StatusOK, jc.Join(userId))
}

//退出社区
var CommunityExit = func(ctx *gin.Context) {
	var param logic.JoinedCommunity
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Exit(userId))
}

//创建社区
var CommunityCreate = func(ctx *gin.Context) {
	var param logic.CommunityCreate
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Create(userId))
}

//修改社区
var CommunityModify = func(ctx *gin.Context) {
	var param logic.CommunityModify
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	userId := ctx.GetInt64("userId")

	//修改logo
	if "" != strings.TrimSpace(param.Logo) {
		ctx.JSON(http.StatusOK, param.Modify(userId, constant.ModifyLogo))
		return
	}

	//修改名字
	if "" != strings.TrimSpace(param.Name) {
		ctx.JSON(http.StatusOK, param.Modify(userId, constant.ModifyName))
		return
	}

	//修改简介
	if "" != strings.TrimSpace(param.Desc) {
		ctx.JSON(http.StatusOK, param.Modify(userId, constant.ModifyDesc))
		return
	}

	ctx.JSON(http.StatusOK, view.SetErr(constant.NoModify))
}

//删除社区
var CommunityDelete = func(ctx *gin.Context) {
	var param logic.CommunityDelete
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	userId := ctx.GetInt64("userId")

	ctx.JSON(http.StatusOK, param.Delete(userId))
}

//增加管理员
var CommunityManagerAdd = func(ctx *gin.Context) {
	var param logic.CommunityManager
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}

	creatorId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Add(creatorId))

}

//删除管理员
var CommunityManagerRemove = func(ctx *gin.Context) {
	var param logic.CommunityManager
	err := ctx.ShouldBindWith(&param, binding.Form)
	if err != nil {
		util.LoggerError(err)
		ctx.JSON(http.StatusOK, view.SetErr(constant.ParamsErr))
		return
	}
	creatorId := ctx.GetInt64("userId")
	ctx.JSON(http.StatusOK, param.Delete(creatorId))
}
