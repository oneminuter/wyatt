package model

//加入的社区
type JoinedCommunity struct {
	TableModel

	UserId        int64  `json:"userId"`
	CommunityId   int64  `json:"communityId"`   //社区id
	CommunityName string `json:"communityName"` //社区名
}
