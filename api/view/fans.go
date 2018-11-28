package view

import "wyatt/api/model"

type Fans struct {
	Account   string `json:"account"`   //账号
	AvatarUrl string `json:"avatarrl"`  //头像
	NickName  string `json:"nickName"`  //昵称
	Sex       int    `json:"sex"`       //性别
	Signature string `json:"signature"` //个性签名
}

func (*Fans) HandlerRespList(muList []model.User) []Fans {
	var list = make([]Fans, 0, len(muList))
	for _, v := range muList {
		list = append(list, Fans{
			Account:   v.Account,
			AvatarUrl: v.AvatarUrl,
			NickName:  v.Nickname,
			Sex:       v.Sex,
			Signature: v.Signature,
		})
	}
	return list
}
