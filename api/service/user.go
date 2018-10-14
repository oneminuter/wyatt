package service

import (
	"errors"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type User struct{}

/*
	构建临时账号
*/
func (*User) GenerateUser(phone, ip string, status int64) model.User {
	account := util.GetRandomString(10, constant.STRING)
	uuid := util.GetUUID()
	ipInfo := util.GetIpInfo(ip)

	return model.User{
		Account:    account,
		Phone:      phone,
		UUID:       uuid,
		NickName:   account,
		AvatarUrl:  constant.DefaultAvator,
		Country:    ipInfo.Data.Country,
		Province:   ipInfo.Data.Region,
		City:       ipInfo.Data.City,
		RegisterIp: ip,
		Status:     status,
	}
}

/*
用户注册，构造用户信息
*/
func (u *User) NewUser(account, pass, ip string) model.User {
	randomStr, password := u.MakePassword(pass)
	ipInfo := util.GetIpInfo(ip)

	return model.User{
		Account:        account,
		Password:       password,
		RandomStr:      randomStr,
		UUID:           util.GetUUID(),
		NickName:       account,
		Sex:            0,
		AvatarUrl:      constant.DefaultAvator,
		Country:        ipInfo.Data.Country,
		Province:       ipInfo.Data.Region,
		City:           ipInfo.Data.City,
		RegisterIp:     ip,
		Status:         1,
		IsSetedAccount: true,
	}
}

/*
	明文加密获得加密后的密码
	明文 MD5 加密之后，追加随机字符串，在进行第二次 MD5 加密得到加密的密码
	返回：追加的随机字符串 和 加密得到的密码
*/
func (*User) MakePassword(str string) (randomStr, password string) {
	p1 := util.MD5(str)
	randomStr = util.GetRandomString(5, constant.STRING)
	password = util.MD5(p1 + randomStr)
	return
}

/*
验证密码
传入加密的随机字符串，明文密码，加密后的密码
*/
func (*User) ValidatePassword(randomStr, pass, password string) bool {
	p1 := util.MD5(pass)
	p2 := util.MD5(p1 + randomStr)
	return p2 == password
}

/*
登录验证
参数： 验证域，域的值，密码
*/
func (u *User) ValidateLogin(field, value, pass string) (model.User, error) {
	var (
		mUser model.User
		err   error
	)
	switch field {
	case "account":
		err = mUser.QueryOne("*", "account = ?", value)
	case "phone":
		err = mUser.QueryOne("*", "phone = ?", value)
	case "email":
		err = mUser.QueryOne("*", "email = ?", value)
	default:
		return model.User{}, errors.New("Illegal field")
	}
	if err != nil {
		util.LoggerError(err)
		return model.User{}, err
	}

	if !u.ValidatePassword(mUser.RandomStr, pass, mUser.Password) {
		return model.User{}, errors.New("Password incorrect")
	}
	return mUser, nil
}

/*
判断账号是否存在
*/
func (*User) IsExitAccount(account string) bool {
	var mUser model.User
	count, err := mUser.QueryCount("account = ?", account)
	if err != nil {
		util.LoggerError(err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

//将用户信息列表转为 id:User 的map
func (*User) TransformToMap(muList []model.User) map[int64]model.User {
	var m = make(map[int64]model.User)
	for _, v := range muList {
		m[v.ID] = v
	}
	return m
}
