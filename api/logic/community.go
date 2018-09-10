package logic

import (
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type Community struct{}

//查询所有状态为 1 的社区
func (c *Community) ListAll() interface{} {
	var m model.Community
	list, err := m.QueryList("*", "status = 1")
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	return list //todo view封装
}
