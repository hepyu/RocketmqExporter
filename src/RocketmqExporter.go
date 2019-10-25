package main

import (
	//"RocketmqExporter/constant"
	"RocketmqExporter/wrapper"
	//"RocketmqExporter/utils"
	//"bytes"
	"fmt"
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
	var content = wrapper.GetConsumerListByTopic("127.0.0.1:30018", "BenchmarkTest")
	fmt.Println("resp:")
	fmt.Println(*content)

}
