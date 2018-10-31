package model

import (
	"time"
	"wyatt/db"

	_ "github.com/jinzhu/gorm"
)

type TableModel struct {
	ID        int64 `json:"-" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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
