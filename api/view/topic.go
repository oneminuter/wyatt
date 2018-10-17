package view

import "wyatt/api/model"

type Topic struct {
	TId              int64  `json:"tId"`                   //话题id，创建的时间戳
	Title            string `json:"title"`                 //标题
	Desc             string `json:"desc" gorm:"type:text"` //简介，详情，或者内容
	CId              int64  `json:"cId"`                   //所属社区id，10位数字
	CreatorAccount   string `json:"creatorAccount"`        //创建者账号
	CreatorAvatarUrl string `json:"creatorAvatarUrl"`      //创建者头像
	CreateTime       int64  `json:"createTime"`            //创建时间戳
	ViewedNum        int64  `json:"viewedNum"`             //浏览量
	ZanNum           int64  `json:"zanNum"`                //点赞量
	CommentNum       int64  `json:"commentNum"`            //评论数量
}

//渲染话题列表
func (*Topic) HandlerRespList(mtList []model.Topic, cId int64, uMap map[int64]model.User) []Topic {
	var list = make([]Topic, 0, len(mtList))
	for _, v := range mtList {
		t := Topic{
			TId:              v.TId,
			Title:            v.Title,
			Desc:             v.Desc,
			CId:              cId,
			CreatorAccount:   uMap[v.CreatorId].Account,
			CreatorAvatarUrl: uMap[v.CreatorId].AvatarUrl,
			CreateTime:       v.CreatedAt.Unix(),
			ViewedNum:        v.ViewedNum,
			ZanNum:           v.ZanNum,
			CommentNum:       v.CommentNum,
		}
		list = append(list, t)
	}
	return list
}

func (t *Topic) HandlerRespDetail(mt model.Topic, cId int64, u model.User) {
	t.TId = mt.TId
	t.Title = mt.Title
	t.Desc = mt.Desc
	t.CId = cId
	t.CreatorAccount = u.Account
	t.CreatorAvatarUrl = u.AvatarUrl
	t.CreateTime = mt.CreatedAt.Unix()
	t.ViewedNum = mt.ViewedNum
	t.ZanNum = mt.ZanNum
	t.CommentNum = mt.CommentNum
}
