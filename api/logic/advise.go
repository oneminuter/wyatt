package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type Advise struct {
	Content string `json:"content" form:"content" binding:"required"`
	Phone   string `json:"phone" form:"phone"`
	Email   string `json:"email" form:"email"`
}
type AdviseList struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

//添加建议
func (a *Advise) Add(userId int64) interface{} {
	ma := model.Advise{
		UserId:  userId,
		Phone:   a.Phone,
		Email:   a.Email,
		Content: a.Content,
	}

	err := ma.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}

//查询用户自个的所提建议列表
func (al *AdviseList) List(userId int64) interface{} {
	var ma model.Advise
	advises, err := ma.QueryList("*", al.Page, al.Limit, "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	var sa view.Advise
	respData := sa.HandlerRespList(advises)
	return view.SetRespData(respData)
}
