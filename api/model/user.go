package model

//用户
type User struct {
	TableModel

	NickName  string `json:"nickName"`  //昵称
	Sex       int    `json:"sex"`       //性别
	Name      string `json:"name"`      //姓名
	AvatarUrl string `json:"avatarUrl"` //头像
	Country   string `json:"country"`   //国家
	Province  string `json:"province"`  //省份
	City      string `json:"city"`      //城市
}
