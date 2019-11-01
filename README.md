笔者第一语言是java，刚开始写go(本程序是笔者的golang处女码)，所以组织风格/写法可能与golang系有比较大的差别。

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

# ().如何编译

有点麻烦，我从开发(IDE用vim)到编译到image制作都是在linux服务器上，所以都是用的golang体系下原生命令进行操作的。

## 1.安装go包依赖管理工具govendor

go get -u -v github.com/kardianos/govendor

## 2.使用govendor下载包依赖

配置环境变量(注意source生效)：export GOPATH=$HOME/go:$HOME/go-workspace

mkdir $HOME/go-workspace/src

然后将本工程clone到目录$HOME/go-workspace/src。

进入$HOME/go-workspace/src执行govendor命令列出工程依赖：govendor list

```
pl  hpy-go-rocketmq-exporter                                   
 l  hpy-go-rocketmq-exporter/constant                          
 l  hpy-go-rocketmq-exporter/model                             
 l  hpy-go-rocketmq-exporter/service                           
 l  hpy-go-rocketmq-exporter/utils                             
 l  hpy-go-rocketmq-exporter/wrapper                           
  m RocketmqExporter/constant                                  
  m RocketmqExporter/model                                     
  m RocketmqExporter/service                                   
  m RocketmqExporter/utils                                     
  m RocketmqExporter/wrapper                                   
  m github.com/go-kit/kit/log/level                            
  m github.com/prometheus/client_golang/prometheus             
  m github.com/prometheus/client_golang/prometheus/promhttp    
  m github.com/prometheus/common/promlog                       
  m github.com/prometheus/common/promlog/flag                  
  m github.com/prometheus/common/version                       
  m gopkg.in/alecthomas/kingpin.v2
```

然后执行govendor init,会生成一个vdendor目录和vendor.json，后边下载的包依赖都会放到这个目录下。

vendor.json

```
{
	"comment": "",
	"ignore": "test",
	"package": [],
	"rootPath": "hpy-go-rocketmq-exporter"
}
```

下载包依赖到vendor目录，执行命令：govendor fetch +out，时间比较长(本工程下提供一个已经编译好的二进制文件：hpy-go-rocketmq-exporter，这个可以直接用于镜像制作)。

执行完成后，vendor目录下：

```
github.com
golang.org
gopkg.in
RocketmqExporter
vendor.json
```

vendor.json内容：

```
{
	"comment": "",
	"ignore": "test",
	"package": [
		{
			"path": "RocketmqExporter/constant",
			"revision": ""
		},
		{
			"path": "RocketmqExporter/model",
			"revision": ""
		},
		{
			"path": "RocketmqExporter/service",
			"revision": ""
		},
		{
			"path": "RocketmqExporter/utils",
			"revision": ""
		},
		{
			"path": "RocketmqExporter/wrapper",
			"revision": ""
		},
		{
			"checksumSHA1": "MXqUZAuWyiMWV7HC0X2krRinZoI=",
			"path": "github.com/alecthomas/template",
			"revision": "fb15b899a75114aa79cc930e33c46b577cc664b1",
			"revisionTime": "2019-07-18T01:26:54Z"
		},
		{
			"checksumSHA1": "3wt0pTXXeS+S93unwhGoLIyGX/Q=",
			"path": "github.com/alecthomas/template/parse",
			"revision": "fb15b899a75114aa79cc930e33c46b577cc664b1",
			"revisionTime": "2019-07-18T01:26:54Z"
		},
		{
			"checksumSHA1": "VT42paM42J+M52CXStvRwsc1v6g=",
			"path": "github.com/alecthomas/units",
			"revision": "f65c72e2690dc4b403c8bd637baf4611cd4c069b",
			"revisionTime": "2019-09-24T02:57:48Z"
		},
		{
			"checksumSHA1": "0rido7hYHQtfq3UJzVT5LClLAWc=",
			"path": "github.com/beorn7/perks/quantile",
			"revision": "37c8de3658fcb183f997c4e13e8337516ab753e6",
			"revisionTime": "2019-07-31T12:00:54Z"
		},
		{
			"path": "github.com/cespare/xxhash/v2",
			"revision": ""
		},
		{
			"checksumSHA1": "eVc+4p1fDrG3e49wZuztY6D2txA=",
			"path": "github.com/go-kit/kit/log",
			"revision": "9f5354e50d79d79d865f684fe139811cf309870f",
			"revisionTime": "2019-10-18T12:22:45Z"
		},
		{
			"checksumSHA1": "dyVQWAYHLspsCzhDwwfQjvkOtMk=",
			"path": "github.com/go-kit/kit/log/level",
			"revision": "9f5354e50d79d79d865f684fe139811cf309870f",
			"revisionTime": "2019-10-18T12:22:45Z"
		},
		{
			"checksumSHA1": "g8yM1TRZyIjXtopiqbslzgLqtM0=",
			"path": "github.com/go-logfmt/logfmt",
			"revision": "07c9b44f60d7ffdfb7d8efe1ad539965737836dc",
			"revisionTime": "2018-11-22T01:56:15Z"
		},
		{
			"checksumSHA1": "Q3FteGbNvRRUMJqbYbmrcBd2DMo=",
			"path": "github.com/golang/protobuf/proto",
			"revision": "ed6926b37a637426117ccab59282c3839528a700",
			"revisionTime": "2019-10-22T19:55:53Z"
		},
		{
			"checksumSHA1": "abKzFXAn0KDr5U+JON1ZgJ2lUtU=",
			"path": "github.com/kr/logfmt",
			"revision": "b84e30acd515aadc4b783ad4ff83aff3299bdfe0",
			"revisionTime": "2014-02-26T03:06:59Z"
		},
		{
			"checksumSHA1": "bKMZjd2wPw13VwoE7mBeSv5djFA=",
			"path": "github.com/matttproud/golang_protobuf_extensions/pbutil",
			"revision": "c182affec369e30f25d3eb8cd8a478dee585ae7d",
			"revisionTime": "2018-12-31T17:19:20Z"
		},
		{
			"checksumSHA1": "I7hloldMJZTqUx6hbVDp5nk9fZQ=",
			"path": "github.com/pkg/errors",
			"revision": "27936f6d90f9c8e1145f11ed52ffffbfdb9e0af7",
			"revisionTime": "2019-02-27T00:00:51Z"
		},
		{
			"checksumSHA1": "HquvlxEmpILGOdePiJzqL/zMvUY=",
			"path": "github.com/prometheus/client_golang/prometheus",
			"revision": "333f01cef0d61f9ef05ada3d94e00e69c8d5cdda",
			"revisionTime": "2019-10-24T23:19:15Z"
		},
		{
			"checksumSHA1": "UBqhkyjCz47+S19MVTigxJ2VjVQ=",
			"path": "github.com/prometheus/client_golang/prometheus/internal",
			"revision": "333f01cef0d61f9ef05ada3d94e00e69c8d5cdda",
			"revisionTime": "2019-10-24T23:19:15Z"
		},
		{
			"checksumSHA1": "UcahVbxaRZ35Wh58lM9AWEbUEts=",
			"path": "github.com/prometheus/client_golang/prometheus/promhttp",
			"revision": "333f01cef0d61f9ef05ada3d94e00e69c8d5cdda",
			"revisionTime": "2019-10-24T23:19:15Z"
		},
		{
			"checksumSHA1": "V8xkqgmP66sq2ZW4QO5wi9a4oZE=",
			"path": "github.com/prometheus/client_model/go",
			"revision": "14fe0d1b01d4d5fc031dd4bec1823bd3ebbe8016",
			"revisionTime": "2019-08-12T15:41:04Z"
		},
		{
			"checksumSHA1": "vA545Z9FkjGvIHBTAKQOE0nap/k=",
			"path": "github.com/prometheus/common/expfmt",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "1Mhfofk+wGZ94M0+Bd98K8imPD4=",
			"path": "github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "ccmMs+h9Jo8kE7izqsUkWShD4d0=",
			"path": "github.com/prometheus/common/model",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "Pj64Wsr2ji1uTv5l49J89Rff0hY=",
			"path": "github.com/prometheus/common/promlog",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "3tSd7cWrq75N2PaoaqAe79Wa+Fw=",
			"path": "github.com/prometheus/common/promlog/flag",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "91KYK0SpvkaMJJA2+BcxbVnyRO0=",
			"path": "github.com/prometheus/common/version",
			"revision": "b5fe7d854c42dc7842e48d1ca58f60feae09d77b",
			"revisionTime": "2019-10-17T12:25:55Z"
		},
		{
			"checksumSHA1": "/otbR/D9hWawJC2jDEqxLdYkryk=",
			"path": "github.com/prometheus/procfs",
			"revision": "34c83637414974b5e7d4bd700b49de3c66631989",
			"revisionTime": "2019-10-22T16:02:49Z"
		},
		{
			"checksumSHA1": "ax1TLBC8m/zLs8u//UHHdFf80q4=",
			"path": "github.com/prometheus/procfs/internal/fs",
			"revision": "34c83637414974b5e7d4bd700b49de3c66631989",
			"revisionTime": "2019-10-22T16:02:49Z"
		},
		{
			"checksumSHA1": "sxRjp2SwHqonjR+sHIEXCkfBglI=",
			"path": "github.com/prometheus/procfs/internal/util",
			"revision": "34c83637414974b5e7d4bd700b49de3c66631989",
			"revisionTime": "2019-10-22T16:02:49Z"
		},
		{
			"path": "golang.org/x/sys/windows",
			"revision": ""
		},
		{
			"checksumSHA1": "sToCp8GThnMnsBzsHv+L/tBYQrQ=",
			"path": "gopkg.in/alecthomas/kingpin.v2",
			"revision": "947dcec5ba9c011838740e680966fd7087a71d0d",
			"revisionTime": "2017-12-17T18:08:21Z"
		}
	],
	"rootPath": "hpy-go-rocketmq-exporter"
}
```

此时make会报错，找不到包github.com/cespare/xxhash/v2，这个是因为prometheus基于依赖于该包，而prometheus是基于gomod构建的，gomod支持能够识别xxhash后面的v2是指定的版本，而redis-shake使用的是govendor不支持版本，解决办法，可以下载github.com/cespare/xxhash/，然后把该文件夹中的内容都copy到github.com/cespare/xxhash/v2目录下即可。

在vendor目录下执行：git clone https://github.com/cespare/xxhash.git github.com/cespare/xxhash/v2

3.编译hpy-go-rocketmq-exporter

执行 make 进行编译，打印信息如下：

```
>> building binaries
GO111MODULE=on /root/go/bin/promu build --prefix /root/go-workspace/src/RocketmqExporter 
 >   RocketmqExporter
>> running all tests
GO111MODULE=on go test -race  -mod=vendor ./...
?   	RocketmqExporter	[no test files]
?   	RocketmqExporter/constant	[no test files]
?   	RocketmqExporter/model	[no test files]
?   	RocketmqExporter/service	[no test files]
?   	RocketmqExporter/utils	[no test files]
?   	RocketmqExporter/wrapper	[no test files]
>> vetting code
GO111MODULE=on go vet  -mod=vendor ./...
```

编译成功后，在目录下会生成一个二进制文件RocketmqExporter，可以直接执行：./RocketmqExporter，打印如下信息说明成功(不用关心报错，因为没有配置参数到环境变量，找不到rocketmq-console)：

```
level=info ts=2019-11-01T09:19:57.879Z caller=RocketmqExporter.go:27 msg="Starting rocketmq_exporter" version="unsupported value type"
level=info ts=2019-11-01T09:19:57.879Z caller=RocketmqExporter.go:28 msg="Build contenxt" (gogo1.13.3,userroot@future,date111911090-09:17:41)=(MISSING)
level=info ts=2019-11-01T09:19:57.879Z caller=RocketmqExporter.go:34 msg=fmt.metricsPath:
panic: http: invalid pattern

goroutine 1 [running]:
net/http.(*ServeMux).Handle(0xd47080, 0x0, 0x0, 0x9f72c0, 0xc000091ec0)
	/usr/local/go/src/net/http/server.go:2397 +0x33a
net/http.Handle(...)
	/usr/local/go/src/net/http/server.go:2446
main.main()
	/root/go-workspace/src/RocketmqExporter/RocketmqExporter.go:39 +0x720
```
