package service

import (
	"RocketmqExporter/constant"
	"RocketmqExporter/model"
	"RocketmqExporter/stringarray"
	"RocketmqExporter/utils"
	"RocketmqExporter/wrapper"
	"fmt"
)

func MsgUnconsumedCount(rocketmqConsoleIPAndPort string) {

	//获取rocketmq集群中的topicNameList
	topicNameArray = wrapper.GetTopicNameList(rocketmqConsoleIPAndPort)
	if topicNameArray == nil {
		return nil
	}

	//获取不纳入监控的topicNameList
	var ignoredTopicNameList = constant.GetIgnoredTopicList()

	for i := range topicNameArray {
		var topicName = topicNameArray[i]
		index = stringarray.contains(constant.GetIgnoredTopicList, topicName)
		if index >= 0 {
			continue
		}

		var data *model.ConsumerList_By_Topic = wrapper.GetConsumerListByTopic(rocketmqConsoleIPAndPort, topicName)

		if data == nil {
			continue
		}

		topicConsumerGroups = data.Data

		for cgName, cgValue := range topicConsumerGroups {
			fmt.Println(topicConsumerGroups[cgName])
		}

	}

}
