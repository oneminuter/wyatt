package handler

import (
	"net/http"
	"wyatt/api/logic"

	"github.com/gin-gonic/gin"
)

//全部社区列表
var CommunityListAll = func(ctx *gin.Context) {
	var lc logic.Community
	ctx.JSON(http.StatusOK, lc.ListAll())
}

//我加入的社区列表
var CommunityListMy = func(ctx *gin.Context) {

}

//加入社区
var Join = func(ctx *gin.Context) {

}

//创建社区
var CommunityCreate = func(ctx *gin.Context) {

}

//修改社区
var CommunityModify = func(ctx *gin.Context) {

}

//删除社区
var CommunityDelete = func(ctx *gin.Context) {

}

//增加管理员
var CommunityManagerAdd = func(ctx *gin.Context) {

}

//删除管理员
var CommunityManagerRemove = func(ctx *gin.Context) {

}

//关注社区
var CommunityFollowAdd = func(ctx *gin.Context) {

}

//取消关注社区
var CommunityFollowRemove = func(ctx *gin.Context) {

}
