package middleware

//请求白名单
var whiteUrlList = []string{
	"/ping",
	"/user/login",
	"/user/register",
	"/advise/add",
	"/user/temp/create",
}

//判断请求url是否在白名单中
func isExitWhite(url string) bool {
	var isExit bool
	for _, v := range whiteUrlList {
		if v == url {
			isExit = true
			break
		}
	}

	if isExit {
		return true
	}
	return false
}
