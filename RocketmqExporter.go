package main

import (
	"RocketmqExporter/constant"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	metricsPath := constant.GetMetricsPath()
	listenAddress := constant.GetListenAddress()
	metricsPrefix := constant.GetMetricsPrefix()

	exporter := DeclareExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
