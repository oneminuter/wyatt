package model

//故事的每条内容
type StoryContent struct {
	TableModel

	StoryId string `json:"storyId"` //属于哪个故事的id
	Type    string `json:"type"`    //1 角色对白，2 旁白
	RoleId  int64  `json:"roleId"`  //角色id, 如果是旁白，该字段为空
	Context string `json:"context"` //内容
}
