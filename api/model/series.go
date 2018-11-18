package model

//系列故事，如小说之类的，有很多篇的，连续的故事
type Series struct {
	TableModel

	Title      string `json:"title"`      //标题，在系列里，相当于书名，在故事章节里相当于每节的标题
	Desc       string `json:"desc"`       //简介
	Classify   string `json:"classify"`   //分类，比如情感类，武侠类，转机类
	AuthorId   int64  `json:"author"`     //作者,对应用户的id
	CoverImg   string `json:"coverImg"`   //本节故事的封面图
	ViewedNum  int64  `json:"viewedNum"`  //浏览量, 此字段在故事章节里是每章节的总浏览量，在系列里是该系列所有故事的总浏览量
	ZanNum     int64  `json:"zanNum"`     //点赞量，此字段在故事章节里是每章节的总点赞量，在系列里是该系列所有故事的总点赞量
	CommentNum int64  `json:"commentNum"` //评论数量，此字段在故事章节里是每章节的总评论量，在系列里是该系列所有故事的总评论量
}
