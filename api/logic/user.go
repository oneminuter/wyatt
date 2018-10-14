package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type User struct{}

//注册参数
type UserRegister struct {
	Account  string `json:"account" form:"account" binding:"required"`   //账号
	Password string `json:"password" form:"password" binding:"required"` //密码
	Ip       string `json:"ip"`                                          //ip
}

//登录参数 account，Phone，Email 三个参数，必须有一个
type UserLogin struct {
	Account  string `json:"account" form:"account"`                      //账号
	Phone    string `json:"phone" form:"phone"`                          //手机号
	Email    string `json:"email" form:"email"`                          //邮箱
	Password string `json:"password" form:"password" binding:"required"` //密码
}

func (*User) Info(userId int64) interface{} {
	//获取用户信息
	var mUser model.User
	err := mUser.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser)
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

/*
注册
*/
func (u *UserRegister) Register() interface{} {
	var sUser service.User

	//判断账号是否存在
	if sUser.IsExitAccount(u.Account) {
		return view.SetErr(constant.AccountIsExisted)
	}

	//构造一个新用户
	mUser := sUser.NewUser(u.Account, u.Password, u.Ip)

	//添加用户到数据库
	err := mUser.Add()
	if err != nil {
		return view.SetErr(constant.CreateUserErr)
	}

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser)
	return view.SetRespData(&vUser)
}

/*
登录
*/
func (u *UserLogin) Login() interface{} {
	var (
		sUser   service.User
		mUser   model.User
		isRight bool
		err     error
	)

	//判断账号是否为空
	if "" != strings.TrimSpace(u.Account) {
		//用账号登录
		mUser, err = sUser.ValidateLogin(`account`, u.Account, u.Password)
		if err == nil {
			isRight = true
		}
	}

	//判断手机号是否为空
	if "" != strings.TrimSpace(u.Phone) && !isRight {
		//验证手机号格式是否错误
		if util.IsPhoneNumber(u.Phone) {
			//用户手机号登录
			mUser, err = sUser.ValidateLogin("phone", u.Phone, u.Password)
			if err == nil {
				isRight = true
			}
		}
	}

	//判断邮箱是否为空
	if "" != strings.TrimSpace(u.Email) && !isRight {
		//验证邮箱格式是否错误
		if util.IsEmail(u.Email) {
			//用邮箱登录
			mUser, err = sUser.ValidateLogin("email", u.Email, u.Password)
			if err == nil {
				isRight = true
			}
		}
	}

	if !isRight {
		//验证失败
		return view.SetErr(constant.LoginErr)
	}

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser)
	return view.SetRespData(&vUser)
}
