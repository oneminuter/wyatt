package view

import (
	"net/http"
	"strings"
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

//mysql查询错误检测
func CheckMysqlErr(err error) interface{} {
	if err != nil && strings.Contains(err.Error(), constant.MysqlNotHaveData) {
		return SetErr(constant.QueryDBEmptyErr)
	}

	//其他非空错误
	return SetErr(constant.QueryDBErr)
}
