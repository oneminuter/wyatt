package logic

import (
	"errors"
	"log"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Comment struct {
	ArticleId string `json:"articleId" form:"articleId" binding:"required"` //所属文章或者话题的流水号
	Page      int    `json:"page" form:"page"`                              //页码，从0开始，默认为0
	Limit     int    `json:"limit" form:"limit"`                            //查询条数, 最大查询20条
}

type CommentAdd struct {
	ArticleId string `json:"articleId" form:"articleId" binding:"required"` //所属文章或者话题的id, 流水号
	Content   string `json:"content" form:"content" binding:"required"`
	ReplyCId  string `json:"replyCid" form:"replyCid"`
}

type CommentDelete struct {
	CId string `json:"articleId" form:"cId" binding:"required"` //评论id, 流水号
}

func (c *Comment) List() interface{} {
	_, _, _, err := util.SplitFlowNumber(c.ArticleId)
	if err != nil {
		return view.SetErr(constant.PasswordIsEmptyErr)
	}

	//查询评论
	var mc model.Comment
	comments, err := mc.QueryList("*", c.Page, c.Limit, "source_flow_id = ?", c.ArticleId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryCommentListErr)
	}

	//提取用户id
	var sc service.Comment
	userIdArr := sc.GetUserIDArr(comments)

	//查询用户信息
	var mu model.User
	userinfoList, err := mu.QueryList("*", 0, c.Limit, "id IN (?)", userIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryCommentListErr)
	}

	//用户信息map, id:user
	var su service.User
	uMap := su.TransformToMap(userinfoList)

	var vc view.Comment
	list := vc.HandlerRespList(comments, uMap)

	return view.SetRespData(list)
}

func (ca *CommentAdd) Add(userId int64) interface{} {
	//判断用户是否禁止在本社区发言
	var mcm model.CommunityManager
	count, err := mcm.QueryCount("user_id = ? AND role = -1", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	if count > 0 {
		return view.SetErr(constant.AccountForbid)
	}

	//分割文章流水号
	tableName, TableID, timestamp, err := util.SplitFlowNumber(ca.ArticleId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	log.Println(tableName, TableID, timestamp)
	//验证流水号是否合法
	if !model.ValidateFlowId(tableName, TableID, timestamp) {
		util.LoggerError(errors.New("流水号不合法"))
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//构造数据，添加
	var mc = model.Comment{
		CreatorId:    userId,
		Content:      ca.Content,
		SourceFlowId: ca.ArticleId,
		ReplyCId:     ca.ReplyCId,
	}
	err = mc.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}

	return view.SetErr(constant.Success)
}

func (cd *CommentDelete) Delete(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(cd.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询评论
	var mc model.Comment
	err = mc.QueryOne("*", "id = ? AND user_id = ?", TableID, userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	err = mc.Delete("id = ?", mc.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}

	return view.SetErr(constant.Success)
}
