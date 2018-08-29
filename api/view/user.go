package view

import (
	"wyatt/api/model"
	"wyatt/util"
)

type User struct {
	Token     string `json:"token"`     //jwt用户标识
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
	u.Token = util.NewToken(mUser.ID, mUser.Status, mUser.UUID)
	u.Account = mUser.Account
	u.UUID = mUser.UUID
	u.NickName = mUser.NickName
	u.Sex = mUser.Sex
	u.Name = mUser.Name
	u.AvatarUrl = mUser.AvatarUrl
	u.Country = mUser.Country
	u.Province = mUser.Province
	u.City = mUser.City
	return
}
