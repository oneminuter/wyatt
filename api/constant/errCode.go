package constant

const (
	Success        = 200
	CreateUserErr  = 600 //创建用户失败
	RequestErr     = 601 //请求错误
	AccountExpire  = 602 //账号过期
	IllegalRequest = 603 //非法请求
	IllegalAccount = 604 //非法账号
	QueryErr       = 605 //查询用户信息失败
)

var ErrMap = map[int]string{
	Success:        "",
	CreateUserErr:  "创建用户失败",
	RequestErr:     "请求错误",
	AccountExpire:  "账号过期",
	IllegalRequest: "非法请求",
	IllegalAccount: "非法账号",
	QueryErr:       "查询失败",
}
