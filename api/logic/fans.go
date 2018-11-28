package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type Fans struct {
	UserAccount string `json:"userAccount" form:"userAccount" binding:"required"` //被关注的用户的账号
}

type FansOrFollowList struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (f *Fans) Follow(userId int64) interface{} {
	//查询被关注用户信息
	var mu model.User
	err := mu.QueryOne("*", "account = ?", f.UserAccount)
	if err != nil {
		return view.CheckMysqlErr(err)
	}

	//判断目标用户状态
	if -1 == mu.Status {
		return view.SetErr(constant.TargetAccountForbid)
	}

	//添加
	var mf = model.Fans{
		UserId:     mu.ID,  //被订阅者的id
		FansUserId: userId, //订阅者的id
	}
	err = mf.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}

func (f *Fans) Cancel(userId int64) interface{} {
	//查询被关注用户信息
	var mu model.User
	err := mu.QueryOne("*", "account = ?", f.UserAccount)
	if err != nil {
		return view.CheckMysqlErr(err)
	}

	//删除
	var mf model.Fans
	err = mf.Delete("user_id = ? AND fans_user_id = ?", mu.ID, userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}

//查询自己的粉丝列表
func (fl *FansOrFollowList) List(userId int64) interface{} {
	//查询粉丝列表
	var mf model.Fans
	fans, err := mf.QueryList("*", fl.Page, fl.Limit, "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取粉丝id数组
	var fansIdArr = make([]int64, len(fans))
	for _, v := range fans {
		fansIdArr = append(fansIdArr, v.FansUserId)
	}

	//查询粉丝信息
	var mu model.User
	ulist, err := mu.QueryList("*", 0, fl.Limit, "id IN (?)", fansIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//返回
	var vf view.Fans
	respData := vf.HandlerRespList(ulist)
	return view.SetRespData(respData)
}

//查询自己的订阅关注列表
func (fl *FansOrFollowList) FollowList(userId int64) interface{} {
	//查询粉丝列表
	var mf model.Fans
	fans, err := mf.QueryList("*", fl.Page, fl.Limit, "fans_user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取粉丝id数组
	var userIdArr = make([]int64, len(fans))
	for _, v := range fans {
		userIdArr = append(userIdArr, v.UserId)
	}

	//查询粉丝信息
	var mu model.User
	ulist, err := mu.QueryList("*", 0, fl.Limit, "id IN (?)", userIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//返回
	var vf view.Fans
	respData := vf.HandlerRespList(ulist)
	return view.SetRespData(respData)
}
