package service

//常用数据初始化
func init() {
	//初始化积分规则数据
	var ir IntegralRule
	ir.cachelRulerIntoMemery()

	//初始化等级规则
	var gl GrowthLevel
	gl.cacheLevelIntoMemery()
}
