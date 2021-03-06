package router

import (
	"wyatt/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("ping", handler.Ping)

	//用户
	userGroup := server.Group("/user")
	userGroup.GET("/info", handler.UserInfo)
	userGroup.POST("/register", handler.UserRegister)
	userGroup.POST("/temp/create", handler.UserTemp)
	userGroup.POST("/login", handler.UserLogin)
	userGroup.POST("/info/modify", handler.UserinfoModify)
	userGroup.POST("/account/modify", handler.UserAccountModify) //修改账号，以为账号目前定的是用户只能修改一次
	userGroup.POST("/password/modify", handler.UserPasswordModify)
	userGroup.POST("/password/reset") //todo

	//社区
	communityGroup := server.Group("/community")
	communityGroup.GET("/list/all", handler.CommunityListAll)
	communityGroup.GET("/list/my", handler.CommunityListMy)
	communityGroup.POST("/join", handler.CommunityJoin)
	communityGroup.POST("/exit", handler.CommunityExit)
	communityGroup.POST("/create", handler.CommunityCreate)
	communityGroup.POST("/modify", handler.CommunityModify)
	communityGroup.POST("/delete", handler.CommunityDelete)
	communityGroup.POST("/manager/add", handler.CommunityManagerAdd)
	communityGroup.POST("/manager/remove", handler.CommunityManagerRemove)

	//话题
	topicGroup := server.Group("/topic")
	topicGroup.GET("/list", handler.TopicList)
	topicGroup.POST("/add", handler.TopicAdd)
	topicGroup.POST("/delete", handler.TopicDelete)
	topicGroup.POST("/modify", handler.TopicModify)
	topicGroup.GET("/detail", handler.TopicDetail)
	topicGroup.POST("/collect/add", handler.TopicCollectAdd)
	topicGroup.POST("/collect/cancel", handler.TopicCollectCancel)
	topicGroup.GET("/collect/list", handler.TopicCollectList)

	//评论
	commentGroup := server.Group("/comment")
	commentGroup.GET("/list", handler.CommentList)
	commentGroup.POST("/add", handler.CommentAdd)
	commentGroup.POST("/delete", handler.CommentDelete)

	//点赞
	zanGroup := server.Group("/zan")
	zanGroup.GET("/list", handler.ZanList)
	zanGroup.POST("/add", handler.ZanAdd)
	zanGroup.POST("/cancel", handler.ZanCancel)

	//粉丝
	fansGroup := server.Group("/fans")
	fansGroup.POST("/follow", handler.FansFollow)
	fansGroup.POST("/cancel", handler.FansCancel)
	fansGroup.GET("/list", handler.FansList)       //粉丝列表
	server.GET("/follow/list", handler.FollowList) //关注列表

	//消息中心
	messageGroup := server.Group("/message")
	messageGroup.GET("/list", handler.MessageList)
	messageGroup.GET("/detail", handler.MessageDetail)
	messageGroup.GET("/delete", handler.MessageDelete)
	messageGroup.POST("/viewed", handler.MessageViewed)

	//用户建议
	adviseGroup := server.Group("/advise")
	adviseGroup.POST("/add", handler.AdviseAdd)
	adviseGroup.GET("/list", handler.AdviseList)

	//举报
	tipGroup := server.Group("/tip")
	tipGroup.POST("/add", handler.TipAdd)

	//故事
	storyGroup := server.Group("/story")
	storyGroup.POST("/add", handler.StoryAdd)
	storyGroup.GET("/list", handler.StoryList) //某用户的所有零散故事列表
	storyGroup.GET("/info", handler.StoryInfo)
	storyGroup.GET("/series/list", handler.StorySeriesList)        //系列列表
	storyGroup.POST("/series/add")                                 //系列添加
	storyGroup.POST("/modify", handler.StoryModify)                //修改标题，简介，封面图之类
	storyGroup.GET("/content/list", handler.StoryContentList)      //故事具体内容列表
	storyGroup.POST("/content/add", handler.StoryContentAdd)       //增加故事细节内容
	storyGroup.POST("/content/modify", handler.StoryContentModify) //修改故事的内容

	//故事角色
	roleGroup := server.Group("/role")
	roleGroup.POST("/add", handler.RoleAdd)
	roleGroup.POST("/modify", handler.RoleModify)
	roleGroup.POST("/delete", handler.RoleDelete)
	roleGroup.GET("/list", handler.RoleList)
	roleGroup.GET("/info", handler.RoleInfo)

	//接口不存在
	server.NoRoute(handler.Page404)
}
