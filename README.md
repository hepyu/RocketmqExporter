笔者第一语言是java，刚开始写go，所以组织风格可能与golang系有些区别。

# ().概述

监控指标：消息堆积数，精确到进程粒度。

监控目的：实时掌控消息消费的健康程度。

数据来源：从rocketmq-console的http请求获取数据。也就是说hpy-go-rocketmq-exporter必须依赖rocketmq-console。好吧，我承认我图省事儿了。。

为什么自己要重新实现： 官方exporter是java的，相对费资源，不适合容器化部署；另外我们要求对消息堆积数有完备监控，且精确到进程级别。
从topic, consumerGroup, broker,queueId, consumerClientIP, consumerClientPID等维度对消息堆积数进行聚合，如下图：

<img src="https://github.com/hepyu/k8s-app-config/blob/master/product/standard/grafana-prometheus-pro/exporter-mq-rocketmq/images/mesage-unconsumed-count.jpg" width="100%">

# ().为何选择golang开发

最适合的选择。常用选型不外乎java, python, golang。

|              语言              |                            优势                             |                       劣势                       |
| ------------------------------------ | ------------------------------------------------------------------- | --------------------------------------------------- |
| java | 写exporter真没啥优势。 | 远比golang和python费资源，容器化下不可接受；相比golang费10倍。|
| python | 比java省资源，但不如golang；开发简单。 | 镜像准备太麻烦；python版本差异太大(我受够了)，不是简单升级个版本就OK的，容器化下python栈可能要维护多批镜像。|
| golang | 开发简单；占用资源很少；性能高。 | 写exporter真没啥劣势。 |

关于镜像大小与实际资源占用的生产对比。

|              语言              |                            K8S生产资源分配                             |                     image大小                         |备注|
| ------------------------------------ | ------------------------------------------------------------------- | --------------------------------------------------- |-------|
| java | cpu:100m, memory:1G。 | 过百兆| 使用官方的rocketmq-exporter，java写的。 |
| python |cpu:100m, memory:100m。|过百兆| 笔者开发，同样依赖于rocketmq-console，位于：https://github.com/hepyu/hpy-rocketmq-exporter |
| golang |cpu:10m, memory:10m。|16MB||

特别说明：

java很不适合开发exporter的重要原因有一点就是，“启动时内存和CPU耗费”与“运行时内存和CPU耗费差异太大”，这就导致容器资源分配时request和max有不小差值，
这个是很不好的，会留下隐患。rocketmq实例不多还好，但是想象一下如果redis,mysql的exporter也是用java写，那这个差值就大了，放大到整个集群将成为潜在风险。
但是如果把request和max设置成一样，又很浪费。

# ().代码组织结构

|              包名               |                            作用                              |                       备注                       |
| ------------------------------------ | ------------------------------------------------------------------- | --------------------------------------------------- |
|constant| 所有的常亮都定义在环境变量中，constant中定义方法取常量| 由于要容器化，舍弃配置文件。|
|model| 存放所有struct结构体，定义要收集的metrics指标。|  |
|utils| 封装工具类，主要是stringarray操作和http访问操作。|
|wrapper| 封账从rocketmq-console或取的数据，并计算汇总成我们要的指标格式。||
|service| 调用wrapper获取数据，计算汇总出消息堆积数的分类统计数据。|根据topic, consumerGroup, broker, clientIP, clientPID等进行分类汇总。|

|              主要代码               |                            作用                              |                       备注                       |
| ------------------------------------ | ------------------------------------------------------------------- | --------------------------------------------------- |
| Collector.go | prometheus的相关代码都在这里，使用prometheus-client将调用service返回的数据写入channel。 ||
| RocketmqExporter.go | 启动http-server，暴露metircs端口。||
