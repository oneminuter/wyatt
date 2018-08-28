package model

//话题
type Topic struct {
	TableModel

	Title       string `json:"title"`       //标题
	Desc        string `json:"desc"`        //简介
	CommunityId int64  `json:"communityId"` //所属社区id
	CreatorId   int64  `json:"creator"`     //创建者id
	ViewedNum   int64  `json:"viewedNum"`   //浏览量
	ZanNum      int64  `json:"zanNum"`      //点赞量
	CommentNum  int64  `json:"commentNum"`  //评论数量
}
