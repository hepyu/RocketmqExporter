package service

import (
	"RocketmqExporter/model"
	"RocketmqExporter/utils"
	"RocketmqExporter/wrapper"
)

func MsgUnconsumedCount(rocketmqConsoleIPAndPort string) {

	topicNameArray = wrapper.GetTopicNameList(rocketmqConsoleIPAndPort)
	if topicNameArray == nil {
		return nil
	}

	for i := range topicNameArray {
		var topicName = topicNameArray[i]
	}

}
