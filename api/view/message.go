package view

import (
	"fmt"
	"wyatt/api/constant"
	"wyatt/api/model"
)

type Message struct {
	MId       string `json:"mId"`                    //消息流水号
	MsgType   string `json:"msgType"`                //消息类型
	Content   string `json:"content"`                //消息内容
	IsViewed  int    `json:"isViewed" gorm:"size:4"` // 是否查看过
	CreatedAt int64  `json:"createdAt"`              //消息产生时间
}

func (*Message) HandleRespList(mList []model.Message) []Message {
	var list = make([]Message, 0, len(mList))
	for _, v := range mList {
		list = append(list, Message{
			MId:       fmt.Sprintf("%s.%d.%d", constant.MG, v.ID, v.FlowId),
			MsgType:   v.MsgType,
			Content:   v.Content,
			IsViewed:  v.IsViewed,
			CreatedAt: v.CreatedAt.Unix(),
		})
	}
	return list
}

func (m *Message) HandleRespDetail(mm model.Message) {
	m.MId = fmt.Sprintf("%s.%d.%d", constant.MG, mm.ID, mm.FlowId)
	m.MsgType = mm.MsgType
	m.Content = mm.Content
	m.IsViewed = mm.IsViewed
	m.CreatedAt = mm.CreatedAt.Unix()
	return
}
