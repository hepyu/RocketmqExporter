package main

import (
	//"RocketmqExporter/constant"
	//"RocketmqExporter/wrapper"
	//"RocketmqExporter/utils"
	"RocketmqExporter/service"
	//"bytes"
	//"fmt"
)

func main() {
	//var test = "hello" + "hpy" + ", how are you?"
	//fmt.Println(test)

	//fmt.Print("rocketmq_exporter")
	//fmt.Print(constant.GetIgnoredTopicList())

	//var content = utils.HttpUrl("http://127.0.0.1:30018/topic/list.query")
	//fmt.Println("content:")
	//fmt.Println(bytes.NewBuffer(content).String())

	//var content = wrapper.GetTopicNameList("127.0.0.1:30018")
	//var content = wrapper.GetConsumerListByTopic("rocketmq-admin.coohua-inc.com", "ne_add_gold_success")
	//fmt.Println("resp:")

	//fmt.Println(content)
	//fmt.Println(*content)

	service.MsgUnconsumedCount("rocketmq-admin.coohua-inc.com")
}
