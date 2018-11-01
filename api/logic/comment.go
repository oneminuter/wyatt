package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Comment struct {
	ArticleId string `json:"articleId" form:"articleId" binding:"required"` //所属文章或者话题的id, 流水号
	Page      int    `json:"page" form:"page"`                              //页码，从0开始，默认为0
	Limit     int    `json:"limit" form:"limit"`                            //查询条数, 最大查询20条
}

type CommentAdd struct {
	ArticleId string `json:"articleId" form:"articleId" binding:"required"` //所属文章或者话题的id, 流水号
	Content   string `json:"content" form:"content" binding:"required"`
	ReplyCId  string `json:"replyCid" form:"replyCid"`
}

func (c *Comment) List() interface{} {
	splits := strings.Split(c.ArticleId, ".")
	if 1 > len(splits) {
		return view.SetErr(constant.QueryDBEmptyErr)
	}
	tableAlias := splits[0]
	AId := splits[1]

	classify, ok := model.TabelMap[tableAlias]
	if !ok {
		return view.SetErr(constant.QueryDBEmptyErr)
	}

	//查询评论
	var mc model.Comment
	comments, err := mc.QueryList("*", c.Page, c.Limit, "classify = ? AND aritcle_id = ?", classify, AId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryCommentListErr)
	}

	//提取用户id
	var sc service.Comment
	userIdArr := sc.GetUserIDArr(comments)

	//查询用户信息
	var mu model.User
	userinfoList, err := mu.QueryList("*", "id IN (?)", userIdArr)
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
	count := mcm.QueryCount("user_id = ? AND role = -1", userId)
	if count > 0 {
		return view.SetErr(constant.AccountForbid)
	}

	//分割文章流水号
	tableName, TableID, timestamp, err := SplitFlowNumber(ca.ArticleId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//验证流水号是否合法
	if !model.ValidateFlowId(tableName, TableID, timestamp) {
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var mc = model.Comment{
		UserId:       userId,
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
