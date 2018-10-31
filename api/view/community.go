package view

import (
	"fmt"
	"wyatt/api/model"
)

type Community struct {
	CreatedAt int64  `json:"createdAt"`  //创建时间
	CId       string `json:"cId"`        //社区号，流水号
	Logo      string `json:"logo"`       //社区logo
	Name      string `json:"name"`       //社区名
	Desc      string `json:"desc"`       //社区描述
	JoinNum   int    `json:"joinNum"`    //加入人数
	TopicNum  int    `json:"articleNum"` //文章数
}

func (c *Community) HandlerRespListAll(mlist []model.Community, joinNumMap, topicNumMap map[int64]int) []Community {
	list := make([]Community, 0)
	for _, v := range mlist {
		c := Community{
			CreatedAt: v.CreatedAt.Unix(),
			CId:       fmt.Sprintf("%s.%d", model.CM),
			Logo:      v.Logo,
			Name:      v.Name,
			Desc:      v.Desc,
			JoinNum:   c.getNum(v.ID, joinNumMap),
			TopicNum:  c.getNum(v.ID, topicNumMap),
		}
		list = append(list, c)
	}
	return list
}

func (*Community) getNum(communityId int64, m map[int64]int) int {
	count, ok := m[communityId]
	if !ok {
		return 0
	}
	return count
}
