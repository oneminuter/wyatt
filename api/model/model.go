package model

import (
	"time"
	"wyatt/db"

	"wyatt/util"

	_ "github.com/jinzhu/gorm"
)

type TableModel struct {
	ID        int64 `json:"-" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FlowId    int64      `json:"flowId" sql:"index"` //子流水号，创建时的时间戳， 10为数字
}

//表别名
const (
	CM   = "CM"   //评论表
	CMT  = "CMT"  //社区表
	JCMT = "JCMT" //加入的社区表
	MG   = "MG"   //消息
	TP   = "TP"   //话题
	Z    = "Z"    //赞
	U    = "U"    //用户
	CMTM = "CMTM" //社区管理员
	TPC  = "TPC"  //话题收藏表
)

//表别名:表名 map
var (
	//表别名:表名 map
	TabelMap = map[string]string{
		CM:   "comments",
		CMT:  "communities",
		JCMT: "joined_communities",
		MG:   "messages",
		TP:   "topics",
		Z:    "zans",
		U:    "users",
		CMTM: "community_managers",
		TPC:  "topic_collects",
	}
)

//最大查询条数

func init() {
	mdb := db.GetMysqlDB()
	mdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Comment{}, &Community{}, &JoinedCommunity{}, &Message{}, &Topic{}, &Zan{}, &User{},
		&CommunityManager{}, &TopicCollect{})
}

//验证流水号是否合法
func ValidateFlowId(tableName string, tableId int64, timestamp int64) bool {
	mdb := db.GetMysqlDB()
	var c int
	err := mdb.Table(tableName).Where("id = ? AND flow_id = ?").Count(&c).Error
	if err != nil {
		util.LoggerError(err)
		return false
	}
	if 1 > c {
		return false
	}
	return true
}
