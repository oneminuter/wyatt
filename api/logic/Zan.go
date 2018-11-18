package logic

import (
	"errors"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Zan struct {
	SourceFlowId string `json:"sourceFlowId" form:"sourceFlowId" binding:"required"`
}
type ZanList struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (za *Zan) Add(userId int64) interface{} {
	//查询是否已经点赞
	var mz model.Zan
	count, err := mz.QueryCount("source_flow_id = ? AND user_id = ?", za.SourceFlowId, userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	if count > 0 {
		return view.SetErr(constant.RepeatOperate)
	}

	//分割流水号
	tableName, TableID, timestamp, err := util.SplitFlowNumber(za.SourceFlowId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询被点赞用户id
	infoMap, err := model.QueryByFlowIdInfo(tableName, TableID, timestamp)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	creatorId, ok := infoMap["creatorId"]
	if !ok {
		util.LoggerError(errors.New("创建者不存在"))
		creatorId = 0
	}

	//添加
	mz = model.Zan{
		UserId:       userId,
		OwnerId:      int64(creatorId.(float64)),
		SourceFlowId: za.SourceFlowId,
	}
	err = mz.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}

	//送积分
	var sir service.IntegralRecord
	go sir.AddIntegral(mz.OwnerId, model.OPT_GET_ZAN)

	return view.SetErr(constant.Success)
}

func (za *Zan) Delete(userId int64) interface{} {
	var mz model.Zan
	err := mz.QueryOne("*", "source_flow_id = ?")
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	if userId != mz.UserId {
		return view.SetErr(constant.ModifyErr)
	}

	err = mz.Delete("id = ?", mz.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}

	//送积分
	var sir service.IntegralRecord
	go sir.AddIntegral(mz.OwnerId, model.OPT_CANCEL_ZAN)

	return view.SetErr(constant.Success)
}

/*
查询用户自己收获的点赞列表
userId: 用户id
*/
func (zl *ZanList) List(userId int64) interface{} {
	// 查询用户的点赞信息
	var mz model.Zan
	zans, err := mz.QueryList("*", zl.Page, zl.Limit, "owner_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取点赞用户id
	var userIdArr = make([]int64, 0, len(zans))
	for _, v := range zans {
		userIdArr = append(userIdArr, v.UserId)
	}

	//查询点赞用户的信息
	var mu model.User
	ulist, err := mu.QueryList("*", 0, zl.Limit, "id IN (?)", userIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	//转换用户信息为map, id:User
	var su service.User
	uMap := su.TransformToMap(ulist)

	//返回
	var vz view.Zan
	respData := vz.HandlerRespList(zans, uMap)

	return view.SetRespData(respData)
}
