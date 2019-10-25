package model

//(1).get topicNameList

type TopicList_Data struct {
	TopicList  []string `json:"topicList"`
	BrokerAddr string   `json:"brokerAddr"`
}

type TopicList struct {
	Status int            `json:"status"`
	Data   TopicList_Data `json:"data"`
	ErrMsg string         `json:"errMsg"`
}

//(2).get consumerList by topic

type ConsumerList_By_Topic struct {
	Status int                   `json:"status"`
	ErrMsg string                `json:"errMsg"`
	Data   map[string]TopicGroup `json:"data"`
}

type TopicGroup struct {
	Topic             string              `json:"topic"`
	DiffTotal         int                 `json:"diffTotal"`
	LastTimestamp     int64               `json:"lastTimestamp"`
	QueueStatInfoList []QueueStatInfoList `json:"queueStatInfoList"`
}

type QueueStatInfoList struct {
	BrokerName     string `json:"brokerName"`
	QueueId        int    `json:"queueId"`
	BrokerOffset   int64  `json:"brokerOffset"`
	ConsumerOffset int64  `json:"consumerOffset"`
	ClientInfo     string `json:"clientInfo"`
	LastTimestamp  int64  `json:"lasttimestamp"`
}
