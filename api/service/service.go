package service

//常用数据初始化
func init() {
	//初始化积分规则数据
	var ir IntegralRuler
	ir.initIntegralRuler()

	//初始化等级规则
	var gl GrowthLevel
	gl.initGrowthLevel()
}
