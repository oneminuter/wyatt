package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
)

type Comment struct {
	Classify  string `json:"classify" form:"classify" binding:"required"`   //所属分类，与表名关联，内部用map对应
	ArticleId string `json:"articleId" form:"articleId" binding:"required"` //所属文章或者话题的id, 表名+创建时的时间戳
	Page      int    `json:"page" form:"page"`                              //页码，从0开始，默认为0
	Limit     int    `json:"limit" form:"limit"`                            //查询条数, 最大查询20条
}

func (c *Comment) List() interface{} {
	table, ok := model.TabelMap[c.Classify]
	if !ok {
		return view.SetErr(constant.QueryDBEmptyErr)
	}

	var mc model.Comment
	mc.QueryList("*", c.Page, c.Limit, "classify = ? AND aritcle_id = ?", table, c.ArticleId)
	return nil
}
