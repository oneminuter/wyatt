package constant

const (
	Success                            = 200
	CreateUserErr                      = 600 //创建用户失败
	RequestErr                         = 601 //请求错误
	AccountExpire                      = 602 //账号过期
	IllegalRequest                     = 603 //非法请求
	IllegalAccount                     = 604 //非法账号
	QueryDBEmptyErr                    = 605 //数据不存在
	ParamsErr                          = 606 //参数错误
	AccountOrPasswordIncludeChinessErr = 607 //账号或者密码不能包含中文字符
	PasswordIsEmptyErr                 = 608 //密码不能为空
	LoginErr                           = 609 //账号或者密码错误
	AccountIsExisted                   = 610 //账号已存在
	AccountForbid                      = 611 //账号已被封禁
	QueryDBErr                         = 612 //查询数据出错
	CommunityProhibition               = 613 //社区已封禁
	CommunityExamining                 = 614 //社区还在申请中
	CommunityDissolution               = 615 //社区已解散
	CommunityJoinErr                   = 616 //加入社区失败
	CommunityCreateErr                 = 617 //创建社区失败
	NoModify                           = 618 //没有修改
	ModifyErr                          = 619 //修改失败
	CommunityIsExist                   = 620 //社区已存在
	NoAuth                             = 621 //没有权限
	DeleteErr                          = 622 //删除失败
	AddErr                             = 623 //添加失败
	UserNotExist                       = 624 //用户不存在
	CommunityExitErr                   = 625 //退出社区失败
	NoCollect                          = 626 //没有收藏
	QueryPageOrLimit                   = 627 //查询页码或者查询条数错误
	CommunityIdErr                     = 628 //社区号错误
	TopicIdErr                         = 629 //话题号错误
	CommentIdErr                       = 630 //评论编号错误
	QueryCommentListErr                = 631 //查询评论列表错误
	IncorrectFlowNumber                = 632 //错误的流水号
	RepeatOperate                      = 633 //重复的操作
	TargetAccountForbid                = 634 //目标用户已被封禁
	MustLogin                          = 635 //你还没有登录，请先登录
	CanModifyOneTime                   = 636 //账号只能修改一次
	TempUserNntCanModify               = 637 //临时用户不能修改
	PasswordErr                        = 638 //密码错误
	UnmarshalContentErr                = 639 //解析内容失败

)

var ErrMap = map[int]string{
	Success:                            "",
	CreateUserErr:                      "创建用户失败",
	RequestErr:                         "请求错误",
	AccountExpire:                      "账号过期",
	IllegalRequest:                     "非法请求",
	IllegalAccount:                     "非法账号",
	QueryDBEmptyErr:                    "数据不存在",
	ParamsErr:                          "参数错误",
	AccountOrPasswordIncludeChinessErr: "账号或者密码不能包含中文字符",
	PasswordIsEmptyErr:                 "密码不能为空",
	LoginErr:                           "账号或者密码错误",
	AccountIsExisted:                   "账号已存在",
	AccountForbid:                      "账号已被封禁",
	QueryDBErr:                         "查询数据出错",
	CommunityProhibition:               "社区已封禁",
	CommunityExamining:                 "社区还在申请中",
	CommunityDissolution:               "社区已解散",
	CommunityJoinErr:                   "加入社区失败",
	CommunityCreateErr:                 "创建社区失败",
	NoModify:                           "没有修改",
	CommunityIsExist:                   "社区已存在",
	NoAuth:                             "没有权限",
	DeleteErr:                          "删除失败",
	AddErr:                             "添加失败",
	UserNotExist:                       "用户不存在",
	CommunityExitErr:                   "退出社区失败",
	NoCollect:                          "没有收藏",
	QueryPageOrLimit:                   "查询页码或者查询条数错误",
	CommunityIdErr:                     "社区号错误",
	TopicIdErr:                         "话题号错误",
	CommentIdErr:                       "评论编号错误",
	QueryCommentListErr:                "查询评论列表错误",
	IncorrectFlowNumber:                "错误的流水号",
	RepeatOperate:                      "重复的操作",
	TargetAccountForbid:                "目标用户已被封禁",
	MustLogin:                          "你还没有登录，请先登录",
	CanModifyOneTime:                   "账号只能修改一次",
	TempUserNntCanModify:               "临时用户不能修改",
	PasswordErr:                        "密码错误",
	UnmarshalContentErr:                "解析内容失败",
}
