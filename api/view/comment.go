package view

import (
	"fmt"
	"wyatt/api/model"
)

type Comment struct {
	CID           string `json:"cid"`           //评论id，流水号
	UserAccount   string `json:"userAccount"`   //评论者账号
	UserAvatarUrl string `json:"userAvatarUrl"` //评论者头像
	CreatedAt     int64  `json:"createdAt"`     //评论时间戳
	Content       string `json:"content"`       //评论内容
	ReplyCID      string `json:"replyCid"`      //回复id, 流水号
}

func (c *Comment) HandlerRespList(mcList []model.Comment, uMap map[int64]model.User) []Comment {
	var list = make([]Comment, 0, len(mcList))
	var (
		u  model.User
		ok bool
	)
	for _, v := range mcList {
		u, ok = uMap[v.UserId]
		if !ok {
			u.Account = ""
			u.AvatarUrl = ""
		}

		list = append(list, Comment{
			CID:           fmt.Sprintf("%s.%d", model.CM, v.CID),
			UserAccount:   u.Account,
			UserAvatarUrl: u.AvatarUrl,
			CreatedAt:     v.CreatedAt.Unix(),
			Content:       v.Content,
			ReplyCID:      v.ReplyCId,
		})
	}
	return list
}
