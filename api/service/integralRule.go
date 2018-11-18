package service

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type IntegralRule struct{}

//缓存积分获取与消费规则
var rules = []model.IntegralRule{}

//缓存积分等级到内存
func (*IntegralRule) cachelRulerIntoMemery() {
	var (
		ir  model.IntegralRule
		err error
	)
	rules, err = ir.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return
	}
	util.Logger("缓存积分或获取与消耗规则")
}

//获取积分规则列表
func (gl *IntegralRule) GetRuleList() []model.IntegralRule {
	if 0 < len(levels) {
		util.Logger("从缓存积分或获取与消耗规则")
		return rules
	}

	//从数据库中获取
	var ir model.IntegralRule
	var err error
	rules, err = ir.QueryList("*", 0, constant.MAX_QUERY_COUNT, "1=1")
	if err != nil {
		util.LoggerError(err)
		return make([]model.IntegralRule, 0)
	}
	return rules
}

//根据操作获取规则明细
func (ir *IntegralRule) GetRule(opt string) model.IntegralRule {
	if 1 > len(rules) {
		ir.GetRuleList()
	}

	for _, v := range rules {
		if opt == v.Operate {
			return v
		}
	}
	return model.IntegralRule{}
}
