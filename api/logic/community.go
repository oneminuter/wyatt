package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Community struct {
	Page  int `json:"page" form:"page"`   //页码，默认从0开始
	Limit int `json:"limit" form:"limit"` //查询条数, 最大查询 constant.MAX_QUERY_COUNT
}
type CommunityCreate struct {
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc" binding:"required"`
}
type CommunityModify struct {
	CId  string `json:"cId" form:"cId" binding:"required"` //表别名+(社区号)创建的时间戳
	Name string `json:"name" form:"name"`                  //社区名
	Desc string `json:"desc" form:"desc"`                  //简介
	Logo string `json:"logo" form:"logo"`                  //logo
}
type CommunityDelete struct {
	CId string `json:"cId" form:"cId" binding:"required"` //表别名+社区号(创建时的时间戳)
}

//查询所有状态为 1 的社区
func (c *Community) ListAll() interface{} {

	var m model.Community
	//获取状态=1(正常)的所有社区
	list, err := m.QueryList("*", c.Page, c.Limit, "status = 1")
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//社区主键id 数组
	var sc service.Community
	communityIdArr := sc.GetCommunityIdArr(list)

	//获取各社区的加入人数
	var mjc model.JoinedCommunity
	joinList, err := mjc.QueryGrounp("community_id, count(*) count", "community_id", "community_id in (?)", communityIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	// 转换为社区id:加入人数map
	joinNumMap := sc.GetCommunityJoinNumMap(joinList)

	//获取社区文章数
	var t model.Topic
	topicList, err := t.QueryGrounp("community_id, count(*) count", "community_id", "community_id in (?)", communityIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	//装换为社区id:文章数map
	topicNumMap := sc.GetCommunityTopicNumMap(topicList)

	var vc view.Community
	resp := vc.HandlerRespListAll(list, joinNumMap, topicNumMap)
	return view.SetRespData(resp)
}

//创建社区
func (cc *CommunityCreate) Create(userId int64) interface{} {
	var c model.Community
	count, err := c.QueryCount("name = ?", cc.Name)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	if count > 0 {
		return view.SetErr(constant.CommunityIsExist)
	}

	c = model.Community{
		Logo:      "",
		Name:      cc.Name,
		Desc:      cc.Desc,
		CreatorId: userId,
		Status:    0,
	}
	err = c.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.CommunityCreateErr)
	}
	return view.SetErr(constant.Success)
}

//修改社区信息
func (cm *CommunityModify) Modify(userId int64, field string) interface{} {
	var (
		c     model.Community
		sc    service.Community
		value string
	)

	_, tableID, _, err := util.SplitFlowNumber(cm.CId)
	if err != nil {
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//判断权限
	isManager := sc.IsManager(tableID, userId)
	if !isManager {
		return view.SetErr(constant.NoAuth)
	}

	switch field {
	case constant.ModifyLogo:
		path, err := sc.SaveLogo(cm.Logo)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.ModifyErr)
		}
		value = path
	case constant.ModifyName:
		value = cm.Name
	case constant.ModifyDesc:
		value = cm.Desc
	default:
		util.Logger("修改选项错误")
		return view.SetErr(constant.ModifyErr)
	}

	err = c.Update(map[string]string{field: value}, "c_id = ?", cm.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}
	return view.SetErr(constant.Success)
}

//删除社区
func (cd *CommunityDelete) Delete(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(cd.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//判断权限
	var sc service.Community
	if !sc.IsAdmin(TableID, userId) {
		return view.SetErr(constant.NoAuth)
	}

	//修改社区的状态为 2
	var mc model.Community
	err = mc.Update(map[string]int{"status": 2}, "c_id = ?", cd.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}
