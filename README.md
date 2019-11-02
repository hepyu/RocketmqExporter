## 微信技术公众号：千里行走

<img src="https://github.com/hepyu/k8s-app-config/blob/master/images/%E5%8D%83%E9%87%8C%E8%A1%8C%E8%B5%B0.jpg" width="25%">

## 实战交流群

<img src="https://github.com/hepyu/saf/blob/master/images/k8s.png" width="25%">

# Summary

监控指标：消息堆积数，精确到进程粒度。

监控目的：实时掌控消息消费的健康程度。

数据来源：从rocketmq-console的http请求获取数据。也就是说RocketmqExporter必须依赖rocketmq-console。好吧，我承认我图省事儿了^_^。

为什么自己要重新实现： 官方exporter是java的，相对费资源；另外我们要求对消息堆积数有完备监控，且精确到进程级别。
从topic, consumerGroup, broker,queueId, consumerClientIP, consumerClientPID等维度对消息堆积数进行聚合，如下图：

<img src="https://github.com/hepyu/k8s-app-config/blob/master/product/standard/grafana-prometheus-pro/exporter-mq-rocketmq/images/mesage-unconsumed-count.jpg" width="100%">

效果图下载地址：https://github.com/hepyu/k8s-app-config/blob/master/product/standard/grafana-prometheus-pro/exporter-mq-rocketmq/images/mesage-unconsumed-count.jpg

# Usage

1.[概述与效果](https://github.com/hepyu/RocketmqExporter/wiki/%E6%A6%82%E8%BF%B0%E4%B8%8E%E6%95%88%E6%9E%9C)

2.[为何选择golang开发](https://github.com/hepyu/RocketmqExporter/wiki/%E4%B8%BA%E4%BD%95%E9%80%89%E6%8B%A9golang%E5%BC%80%E5%8F%91) (附带不同语言开发的优劣对比)

3.[代码组织结构与文件说明](https://github.com/hepyu/RocketmqExporter/wiki/%E4%BB%A3%E7%A0%81%E7%BB%84%E7%BB%87%E7%BB%93%E6%9E%84%E4%B8%8E%E6%96%87%E4%BB%B6%E8%AF%B4%E6%98%8E)

4.[如何编译](https://github.com/hepyu/RocketmqExporter/wiki/%E5%A6%82%E4%BD%95%E7%BC%96%E8%AF%91)

5.[相关编译文件说明](https://github.com/hepyu/RocketmqExporter/wiki/%E7%9B%B8%E5%85%B3%E7%BC%96%E8%AF%91%E6%96%87%E4%BB%B6%E8%AF%B4%E6%98%8E)

6.[如何进行容器化部署](https://github.com/hepyu/RocketmqExporter/wiki/%E5%A6%82%E4%BD%95%E8%BF%9B%E8%A1%8C%E5%AE%B9%E5%99%A8%E5%8C%96%E9%83%A8%E7%BD%B2)

7.[如何进行实体机部署](https://github.com/hepyu/RocketmqExporter/wiki/%E5%A6%82%E4%BD%95%E8%BF%9B%E8%A1%8C%E5%AE%9E%E4%BD%93%E6%9C%BA%E9%83%A8%E7%BD%B2)

8.[如何结合prometheus与grafana
](https://github.com/hepyu/RocketmqExporter/wiki/%E5%A6%82%E4%BD%95%E7%BB%93%E5%90%88prometheus%E4%B8%8Egrafana)

# TODO

后续有时间会把官方rocketmq-exporter的监控指标也用go重写。
