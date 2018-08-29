package model

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
