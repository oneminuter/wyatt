package model

//故事 - 故事可多人参与创造
type Story struct {
	Series

	SeriesId int64 `json:"seriesId"` //所属系列id, 如有章节的小说id, 此字段的有无由于区分是零散的故事，还是有章节、连续性的的故事或者小说
}
