package service

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type User struct{}

func (*User) GenerateUser(phone, ip string, status int64) model.User {
	account := util.GetRandomString(10, constant.STRING)
	uuid := util.GetUUID()
	ipInfo := util.GetIpInfo(ip)

	return model.User{
		Account:    account,
		Phone:      phone,
		UUID:       uuid,
		NickName:   account,
		AvatarUrl:  "http://blog.oneminuter.com/favicon.ico",
		Country:    ipInfo.Data.Country,
		Province:   ipInfo.Data.Region,
		City:       ipInfo.Data.City,
		RegisterIp: ip,
		Status:     status,
	}
}
