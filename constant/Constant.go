package constant

import (
	"os"
	"strings"
)

//rockemq有一些默认topic，这些不需要纳入监控，需要支持可配置
//var ignoredTopicList = []string{"RMQ_SYS_TRANS_HALF_TOPIC", "BenchmarkTest", "OFFSET_MOVED_EVENT", "TBW102", "SELF_TEST_TOPIC", "DefaultCluster", "broker-b", "broker-a"}
func GetIgnoredTopicArray() []string {
	var ignoredTopicsInEnv string = os.Getenv("ignoredTopics")
	var tarray []string = strings.Split(ignoredTopicsInEnv, ",")
	return tarray
}

//配置要监控的rocketmq集群，从rocketmqconsole获取数据
func GetRocketmqConsoleIPAndPort() string {
	return os.Getenv("rocketmqConsoleIPAndPort")
}

//定义metrics接口path
func GetMetricsPath() string {
	return os.Getenv("metricsPath")
}

//定义暴露的端口
func GetListenAddress() string {
	return os.Getenv("listenAddress")
}

//定义metrics数据名称的前缀：比如定义前缀是rocketmq，则metrics为:rocketmq_msg_diff_detail等
func GetMetricsPrefix() string {
	return os.Getenv("metricsPrefix")
}
