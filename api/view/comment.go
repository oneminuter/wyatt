package view

type Comment struct {
	CID           int64  `json:"cid"`           //评论id，10位数字
	UserAccount   string `json:"userAccount"`   //评论者账号
	UserAvatarUrl string `json:"userAvatarUrl"` //评论者头像
	CreatedAt     int64  `json:"createdAt"`     //评论时间戳
	Content       string `json:"content"`       //评论内容
	ReplyCID      int64  `json:"replyCid"`      //回复id
}
