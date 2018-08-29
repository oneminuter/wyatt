package view

import "wyatt/api/model"

type User struct {
	Account   string `json:"account"`   //账号
	UUID      string `json:"uuid"`      //用户标识
	NickName  string `json:"nickName"`  //昵称
	Sex       int    `json:"sex"`       //性别
	Name      string `json:"name"`      //姓名
	AvatarUrl string `json:"avatarUrl"` //头像
	Country   string `json:"country"`   //国家
	Province  string `json:"province"`  //省份
	City      string `json:"city"`      //城市
}

func (u *User) RenderUserInfo(mUser *model.User) {
	u = &User{
		Account:   mUser.Account,
		UUID:      mUser.UUID,
		NickName:  mUser.NickName,
		Sex:       mUser.Sex,
		Name:      mUser.Name,
		AvatarUrl: mUser.AvatarUrl,
		Country:   mUser.Country,
		Province:  mUser.Province,
		City:      mUser.City,
	}
}
