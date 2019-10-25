package constant

//var ignoredTopicList = []string{"RMQ_SYS_TRANS_HALF_TOPIC", "BenchmarkTest", "OFFSET_MOVED_EVENT", "TBW102", "SELF_TEST_TOPIC", "DefaultCluster", "broker-b", "broker-a"}

var ignoredTopicList = []string{}

func GetIgnoredTopicList() []string {
	return ignoredTopicList
}
