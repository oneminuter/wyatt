package view

import (
	"log"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type Zan struct {
	UserAccount   string `json:"userAccount"`   //点赞用户的账号
	UserAvatarUrl string `json:"userAvatarUrl"` //点击用户的头像
	SourceFlowId  string `json:"sourceFlowId"`  //被点赞的资源流水号
	Classify      string `json:"classify"`      //被点赞资源的类型分类，如文章，话题，个人资料
	CreatedAt     int64  `json:"createdAt"`     //点赞时间戳
}

func (*Zan) HandlerRespList(zans []model.Zan, uMap map[int64]model.User) []Zan {
	var (
		list     = make([]Zan, 0, len(zans))
		u        model.User
		ok       bool
		classify string
	)
	for _, v := range zans {
		u, ok = uMap[v.UserId]
		if !ok {
			u.Account = constant.DefaultAvator
			u.AvatarUrl = ""
		}

		//根据表别名取分类
		tableName, _, _, _ := util.SplitFlowNumber(v.SourceFlowId)
		log.Println(tableName)
		classify, ok = constant.TableAliasNameMap[tableName]
		if !ok {
			classify = ""
		}

		list = append(list, Zan{
			UserAccount:   u.Account,
			UserAvatarUrl: u.AvatarUrl,
			SourceFlowId:  v.SourceFlowId,
			Classify:      classify,
			CreatedAt:     v.CreatedAt.Unix(),
		})
	}
	return list
}
