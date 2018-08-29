package util

import (
	"encoding/json"
	"wyatt/api/model"

	"github.com/levigross/grequests"
)

//获取ip信息
func GetIpInfo(ip string) model.Location {
	resp, err := grequests.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+ip, nil)
	if err != nil {
		LoggerError(err)
	}

	var loc model.Location
	err = json.Unmarshal(resp.Bytes(), &loc)
	if err != nil {
		LoggerError(err)
	}

	return loc
}
