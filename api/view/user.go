package view

import (
	"wyatt/api/model"
	"wyatt/util"
)

type User struct {
	Token     string `json:"token"`     //jwt用户标识
	Account   string `json:"account"`   //账号
	UUID      string `json:"uuid"`      //用户标识
	Email     string `json:"email"`     //邮箱
	Nickname  string `json:"nickname"`  //昵称
	Sex       int    `json:"sex"`       //性别
	Name      string `json:"name"`      //姓名
	AvatarUrl string `json:"avatarUrl"` //头像
	Signature string `json:"signature"` //个性签名
	Country   string `json:"country"`   //国家
	Province  string `json:"province"`  //省份
	City      string `json:"city"`      //城市
	Avaliable int    `json:"avaliable"` //可用积分，可消耗
	Level     int    `json:"level"`     //用户等级
}

func (u *User) HandlerRespUserInfo(mUser *model.User, mi *model.Integral) {
	u.Token = util.NewToken(mUser.ID, mUser.Status, mUser.UUID)
	u.Account = mUser.Account
	u.UUID = mUser.UUID
	u.Email = mUser.Email
	u.Nickname = mUser.Nickname
	u.Sex = mUser.Sex
	u.Name = mUser.Name
	u.AvatarUrl = mUser.AvatarUrl
	u.Signature = mUser.Signature
	u.Country = mUser.Country
	u.Province = mUser.Province
	u.City = mUser.City
	u.Avaliable = mi.Avaliable
	u.Level = mi.Level
	return
}
