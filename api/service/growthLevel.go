package service

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type GrowthLevel struct{}

//启动时，将等级划分规则加载到内存中
var levels = []model.GrowthlLevel{}

func (*GrowthLevel) initGrowthLevel() {
	var (
		mgl model.GrowthlLevel
		err error
	)
	levels, err = mgl.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return
	}
	util.Logger("缓存积分等级规则")
}

//获取等级划分列表
func (gl *GrowthLevel) GetGrowthLevelList() []model.GrowthlLevel {
	if 0 < len(levels) {
		util.Logger("从缓存中取积分等级规则")
		return levels
	}

	//从数据库中获取
	var mgl model.GrowthlLevel
	var err error
	levels, err = mgl.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return make([]model.GrowthlLevel, 0)
	}
	return levels
}

//根据成长值获取等级
func (gl *GrowthLevel) GetGrowthLevel(growth int) int {
	if 1 > len(levels) {
		gl.GetGrowthLevelList()
	}

	//倒叙查找
	for i := len(levels) - 1; i > 0; i-- {
		if growth >= levels[i].Start {
			return levels[i].Level
		}
	}
	return 0
}
