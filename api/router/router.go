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
	userGroup.POST("/info/modify")
	userGroup.POST("/register", handler.UserRegister)
	userGroup.POST("/login", handler.UserLogin)
	userGroup.POST("/password/modify")
	userGroup.POST("/password/reset")

	//社区
	communityGroup := server.Group("/community")
	communityGroup.GET("/list/all", handler.CommunityListAll)
	communityGroup.GET("/list/my", handler.CommunityListMy)
	communityGroup.POST("/join", handler.Join)
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
	commentGroup.GET("/list")
	commentGroup.POST("/add")
	commentGroup.POST("/delete")

	//点赞
	zanGroup := server.Group("/zan")
	zanGroup.GET("/list")
	zanGroup.POST("/add")
	zanGroup.POST("/cancel")

	//粉丝
	fansGroup := server.Group("/fans")
	fansGroup.GET("/list")
	fansGroup.POST("/follow")
	fansGroup.POST("/cancel")

	//消息中心
	messageGroup := server.Group("/message")
	messageGroup.GET("/list")
	messageGroup.GET("/detail")
	messageGroup.GET("/delete")
}
