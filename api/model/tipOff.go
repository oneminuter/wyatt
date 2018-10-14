package model

//内容举报
type TipOff struct {
	TableModel
	UserId   int64  `json:"userId"`   //用户id
	Classify string `json:"classify"` //分类，目前定义是和表名关联的字段，用于定位被举报的内容所属表
	ItemId   int64  `json:"itemId"`   //被举报内容的公有id, 比如话题的TId, 社区的CId, 为创建条目是时间戳
	Status   int    `json:"status"`   //处理状态
	Remark   string `json:"remark"`   //处理备注
}
