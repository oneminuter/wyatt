package service

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type IntegralRuler struct{}

//缓存积分获取与消费规则
var rulers = []model.IntegralRuler{}

func init() {
	var (
		ir  model.IntegralRuler
		err error
	)
	rulers, err = ir.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return
	}
	util.Logger("缓存积分或获取与消耗规则")
}

//获取积分列表
func (gl *IntegralRuler) GetIntegralRulerList() []model.IntegralRuler {
	if 0 < len(levels) {
		util.Logger("从缓存积分或获取与消耗规则")
		return rulers
	}

	//从数据库中获取
	var ir model.IntegralRuler
	var err error
	rulers, err = ir.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return make([]model.IntegralRuler, 0)
	}
	return rulers
}

//根据操作获取规则明细
func (ir *IntegralRuler) GetGrowth(opt string) model.IntegralRuler {
	if 1 > len(rulers) {
		ir.GetIntegralRulerList()
	}

	for _, v := range rulers {
		if opt == v.Operate {
			return v
		}
	}
	return model.IntegralRuler{}
}
