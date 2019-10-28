package wrapper

import (
	"RocketmqExporter/model"
	"RocketmqExporter/utils"
	//"bytes"
	"encoding/json"
	"fmt"
)

func GetTopicNameList(rocketmqConsoleIPAndPort string) []string {
	var url = "http://" + rocketmqConsoleIPAndPort + "/topic/list.query"
	var content = utils.HttpUrl(url)

	var jsonData model.TopicList
	err := json.Unmarshal([]byte(content), &jsonData)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return jsonData.Data.TopicList
}

func GetConsumerListByTopic(rocketmqConsoleIPAndPort string, topicName string) *model.ConsumerList_By_Topic {
	/*
			var content = `{
				"status": 0,
				"data": {
					"topic1_success_for_ma": {
						"topic": "topic1_success",
						"diffTotal": 8,
						"lastTimestamp": 1571911971666,
						"queueStatInfoList": [{
							"brokerName": "broker-a",
							"queueId": 0,
							"brokerOffset": 2538968172,
							"consumerOffset": 2538968171,
							"clientInfo": "ip1@4979",
							"lastTimestamp": 1571911971656
						}]
					},
					"topic1_success_for_task": {
						"topic": "topic1_success",
						"diffTotal": 8,
						"lastTimestamp": 1571911971674,
						"queueStatInfoList": [{
							"brokerName": "broker-a",
							"queueId": 0,
							"brokerOffset": 2538968173,
							"consumerOffset": 2538968172,
							"clientInfo": "ip2@7219",
							"lastTimestamp": 1571911971660
						}]
					},
					"topic1_success_for_activity": {
						"topic": "topic1_success",
						"diffTotal": 7,
						"lastTimestamp": 1571911971683,
						"queueStatInfoList": [{
							"brokerName": "broker-a",
							"queueId": 0,
							"brokerOffset": 2538968175,
							"consumerOffset": 2538968174,
							"clientInfo": "ip3@16015",
							"lastTimestamp": 1571911971679
						}]
					}
				},
				"errMsg": null
			}`

		var jsonData model.ConsumerList_By_Topic
		err := json.Unmarshal([]byte(content), &jsonData)
		fmt.Println(err)
		fmt.Printf("%+v", jsonData)
	*/

	var url = "http://" + rocketmqConsoleIPAndPort + "/topic/queryConsumerByTopic.query?topic=" + topicName
	var content = utils.HttpUrl(url)

	var jsonData *model.ConsumerList_By_Topic
	err := json.Unmarshal([]byte(content), &jsonData)

	if err != nil {
		return nil
	}

	return jsonData
}
