package service

import "wyatt/api/model"

type TopicCollect struct{}

func (TopicCollect) GetTopicIdArr(collects []model.TopicCollect) []int64 {
	var topicIdArr = make([]int64, 0, len(collects))
	for _, c := range collects {
		topicIdArr = append(topicIdArr, c.TopicId)
	}
	return topicIdArr
}
