package constant

const (
	STRING        = "string"
	NUMBER        = "number"
	TokenExpired  = "token is expired"
	DefaultAvator = "http://blog.oneminuter.com/favicon.ico"
	//数据为空
	RedisNotHaveData = "redis: nil"
	MysqlNotHaveData = "record not found"

	//社区修改 - 字段
	ModifyLogo = "logo" //标题
	ModifyName = "name" //名字
	ModifyDesc = "desc" //简介

	//最大查询条数
	MAX_QUERY_COUNT = 100

	//表别名
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
	TableAliasNameMap = map[string]string{
		TabelMap[CM]:   "评论",
		TabelMap[CMT]:  "社区",
		TabelMap[JCMT]: "加入的社区",
		TabelMap[MG]:   "消息",
		TabelMap[TP]:   "话题",
		TabelMap[Z]:    "点赞",
		TabelMap[U]:    "用户",
		TabelMap[CMTM]: "社区管理",
		TabelMap[TPC]:  "话题收藏",
	}
)
