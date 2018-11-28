package service

import (
	"wyatt/api/model"
	"wyatt/util"
)

type IntegralRecord struct{}

//获得积分
func (*IntegralRecord) AddIntegral(userId int64, operate string) error {
	var ir IntegralRule
	rule := ir.GetRule(operate)

	growth := float64(rule.Integral) * rule.SpeedRate

	var mir = model.IntegralRecord{
		UserId: userId,
		Growth: int(growth),
	}
	mir.Operate = rule.Operate
	mir.Integral = rule.Integral
	mir.SpeedRate = rule.SpeedRate

	err := mir.Add()
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}
