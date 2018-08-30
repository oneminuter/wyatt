package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type User struct{}

//注册参数，account,phone,email 三个参数，必须有一个
type UserRegister struct {
	Account  string `json:"account"`                                     //账号
	Phone    string `json:"phone"`                                       //手机号
	Email    string `json:"email"`                                       //邮箱
	Password string `json:"password" form:"password" binding:"required"` //密码
	NickName string `json:"nickName"`                                    //昵称
}

func (*User) Info(userId int64) interface{} {
	//获取用户信息
	var mUser model.User
	err := mUser.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryErr)
	}

	//返回数据
	var vUser view.User
	vUser.RenderUserInfo(&mUser)

	return view.SetRespData(&vUser)
}

//新增临时用户
func (*User) AddTempUser(ip string) (userId int64, err error) {
	//构建一个用户
	var sUser service.User
	mUser := sUser.GenerateUser("", ip, 0)

	//存入数据库
	err = mUser.Add()
	if err != nil {
		util.LoggerError(err)
	}
	userId = mUser.ID
	return
}
