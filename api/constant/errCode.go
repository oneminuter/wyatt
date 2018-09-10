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
	LoginErr                           = 609 //账号或则密码错误
	AccountIsExisted                   = 610 //账号已存在
	AccountForbid                      = 611 //账号已被封禁
	QueryDBErr                         = 612 //查询数据出错
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
	LoginErr:                           "账号或则密码错误",
	AccountIsExisted:                   "账号已存在",
	AccountForbid:                      "账号已被封禁",
	QueryDBErr:                         "查询数据出错",
}
