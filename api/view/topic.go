package view

import (
	"fmt"
	"wyatt/api/model"
)

type Topic struct {
	TId              string `json:"tId"`                   //话题id，流水号
	Title            string `json:"title"`                 //标题
	Desc             string `json:"desc" gorm:"type:text"` //简介，详情，或者内容
	CId              string `json:"cId"`                   //所属社区id，流水号
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
			TId:              fmt.Sprintf("%s.%d", model.TP, v.TId),
			Title:            v.Title,
			Desc:             v.Desc,
			CId:              fmt.Sprintf("%s.%d", model.CMT, cId),
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

//话题详情
func (t *Topic) HandlerRespDetail(mt model.Topic, cId int64, u model.User) {
	t.TId = fmt.Sprintf("%s.%d", model.TP, mt.TId)
	t.Title = mt.Title
	t.Desc = mt.Desc
	t.CId = fmt.Sprintf("%s.%d", model.CMT, cId)
	t.CreatorAccount = u.Account
	t.CreatorAvatarUrl = u.AvatarUrl
	t.CreateTime = mt.CreatedAt.Unix()
	t.ViewedNum = mt.ViewedNum
	t.ZanNum = mt.ZanNum
	t.CommentNum = mt.CommentNum
}

//收藏话题列表
func (*Topic) HandlerRespCollectList(mtList []model.Topic, cIDMap map[int64]int64, uMap map[int64]model.User) []Topic {
	var (
		ok        bool
		cId       int64 //社区id - 10为数字
		u         model.User
		avatarUrl string
		account   string
		list      = make([]Topic, 0, len(mtList))
	)
	for _, v := range mtList {
		cId, ok = cIDMap[v.CommunityId]
		if !ok {
			cId = 0
		}

		u, ok = uMap[v.CreatorId]
		if ok {
			avatarUrl = u.AvatarUrl
			account = u.Account
		} else {
			avatarUrl = ""
			account = ""
		}

		t := Topic{
			TId:              fmt.Sprintf("%s.%d", model.TP, v.TId),
			Title:            v.Title,
			Desc:             v.Desc,
			CId:              fmt.Sprintf("%s.%d", model.CMT, cId),
			CreatorAccount:   account,
			CreatorAvatarUrl: avatarUrl,
			CreateTime:       v.CreatedAt.Unix(),
			ViewedNum:        v.ViewedNum,
			ZanNum:           v.ZanNum,
			CommentNum:       v.CommentNum,
		}
		list = append(list, t)
	}
	return list
}
