package view

import "wyatt/api/model"

type Advise struct {
	Content   string `json:"content"` //建议内容
	Status    int    `json:"status" ` //处理状态， 默认 0, 1 已看
	Remark    string `json:"remark"`  //备注
	CreatedAt int64  `json:"createdAt"`
}

func (a *Advise) HandlerRespList(mlist []model.Advise) []Advise {
	var list = make([]Advise, 0, len(mlist))
	for _, v := range mlist {
		list = append(list, Advise{
			Content:   v.Content,
			Status:    v.Status,
			Remark:    v.Remark,
			CreatedAt: v.CreatedAt.Unix(),
		})
	}
	return list
}
