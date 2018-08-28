package view

type response struct {
	ErrNum int         `json:"errNum"` //错误码
	ErrMsg string      `json:"errMsg"` //错误信息
	Data   interface{} `json:"data"`   //返回数据
}
