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

//修改用户信息
type UserinfoModify struct {
	Phone     string `json:"phone" form:"phone"`         //电话
	Nickname  string `json:"nickName" form:"nickname"`   //昵称
	Sex       int    `json:"sex" form:"sex"`             //性别
	Name      string `json:"name" form:"name"`           //姓名
	Email     string `json:"email" form:"email"`         //邮箱
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"` //头像
	Signature string `json:"signature" form:"signature"` //个性签名
}

type UserAccountModify struct {
	Account string `json:"account" form:"account" binding:"required"` //账号
}

//用户信息
func (*User) Info(userId int64) interface{} {
	//获取用户信息
	var mUser model.User
	err := mUser.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//查询积分信息
	var mi model.Integral
	mi.UserId = userId
	err = mi.QueryOne("*", "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser, &mi)
	return view.SetRespData(&vUser)
}

//新增临时用户
func (*User) AddTempUser(ip string) interface{} {
	//构建一个用户
	var sUser service.User
	mUser := sUser.GenerateUser("", ip, 0)

	//存入数据库
	err := mUser.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.CreateUserErr)
	}

	//返回
	var vu view.User
	vu.HandlerRespUserInfo(&mUser, &model.Integral{})
	return view.SetRespData(vu)
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

	//积分初始化
	var mi = model.Integral{
		Avaliable: 0,
		Growth:    0,
		UserId:    mUser.ID,
		Level:     0,
	}
	err = mi.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}

	//送积分
	var sir service.IntegralRecord
	go sir.AddIntegral(mUser.ID, model.OPT_REGISTER)

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser, &mi)
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

	var mi model.Integral
	err = mi.QueryOne("*", "user_id = ?", mUser.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//送积分
	var sir service.IntegralRecord
	go sir.AddIntegral(mUser.ID, model.OPT_LOGIN)

	//返回数据
	var vUser view.User
	vUser.HandlerRespUserInfo(&mUser, &mi)
	return view.SetRespData(&vUser)
}

//修改用户信息
func (um *UserinfoModify) Modify(userId int64) interface{} {
	var modify = make(map[string]interface{})
	if "" != strings.TrimSpace(um.Name) {
		modify["name"] = um.Name
	}
	if "" != strings.TrimSpace(um.AvatarUrl) {
		modify["avatar_url"] = um.AvatarUrl
	}
	if "" != strings.TrimSpace(um.Email) {
		modify["email"] = um.Email
	}
	if "" != strings.TrimSpace(um.Phone) {
		modify["phone"] = um.Phone
	}
	if "" != strings.TrimSpace(um.Signature) {
		modify["signature"] = um.Signature
	}
	if "" != strings.TrimSpace(um.Nickname) {
		modify["nick_name"] = um.Nickname
	}
	if 0 < um.Sex && 3 > um.Sex {
		modify["sex"] = um.Sex
	}

	//修改
	var mu model.User
	err := mu.Update(modify, "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}
	return view.SetErr(constant.Success)
}

//用户修改账号
func (uam *UserAccountModify) Modify(userId int64) interface{} {
	//查询用户信息
	var mu model.User
	err := mu.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//判断是否是临时用户，临时用户不能修改账号
	if 0 == mu.Status {
		return view.SetErr(constant.TempUserNntCanModify)
	}

	//判断用户是否第一次修改
	if mu.IsSetedAccount {
		return view.SetErr(constant.CanModifyOneTime)
	}

	//修改
	mu.Account = uam.Account
	err = mu.Update(map[string]string{"account": uam.Account}, "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}
	return view.SetErr(constant.Success)
}
