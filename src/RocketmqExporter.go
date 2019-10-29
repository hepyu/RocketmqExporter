package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func metrics() {

}

func main() {

	metricsPath := "/metrics"
	listenAddress := ":8081"
	metricsPrefix := "rocketmq"

	exporter := DeclareExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
