package util

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type Location struct {
	Code int  `json:"code"`
	Data data `json:"data"`
}

type data struct {
	Ip        string `json:"ip"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	Region    string `json:"region"`
	City      string `json:"city"`
	County    string `json:"county"`
	Isp       string `json:"isp"`
	CountryId string `json:"country_id"`
	AreaId    string `json:"area_id"`
	RegionId  string `json:"region_id"`
	CityId    string `json:"city_id"`
	CountyId  string `json:"county_id"`
	IspId     string `json:"isp_id"`
}

//获取ip信息
func GetIpInfo(ip string) Location {
	resp, err := grequests.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+ip, nil)
	if err != nil {
		LoggerError(err)
	}

	var loc Location
	err = json.Unmarshal(resp.Bytes(), &loc)
	if err != nil {
		LoggerError(err)
	}

	return loc
}
