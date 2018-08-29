package view

import (
	"net/http"
	"wyatt/api/constant"
)

type Response struct {
	ErrCode int         `json:"errCode"` //错误码
	ErrMsg  string      `json:"errMsg"`  //错误信息
	Data    interface{} `json:"data"`    //返回数据
}

func SetErr(code int) Response {
	return Response{
		ErrCode: code,
		ErrMsg:  constant.ErrMap[code],
	}
}

func SetRespData(data interface{}) Response {
	return Response{
		ErrCode: http.StatusOK,
		ErrMsg:  constant.ErrMap[constant.Success],
		Data:    data,
	}
}
