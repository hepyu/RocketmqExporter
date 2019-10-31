package main

import (
	"RocketmqExporter/constant"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

func main() {

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)

	kingpin.Version(version.Print("rocketmq_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting rocketmq_exporter", "version", version.Info)
	level.Info(logger).Log("msg", "Build contenxt", version.BuildContext())

	metricsPath := constant.GetMetricsPath()
	listenAddress := constant.GetListenAddress()
	metricsPrefix := constant.GetMetricsPrefix()

	level.Info(logger).Log("msg", "fmt.metricsPath:"+metricsPath)

	exporter := DeclareExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
