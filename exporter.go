package main

import (
	"flag"
	"fmt"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
	"nginxexporter/pkg/collector/topuri"
	"nginxexporter/pkg/crcache"
)

const (
	metricUrl = "/metrics"
)

var (
	ListenAddress = flag.String("listen-address", "0.0.0.0:8080", "the address of exporter binding")
)


func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})
	crCache := crcache.NewCrCache()
	crCache.Add("test1", 1)
	crCache.Add("test2", 1)

	topUriMetricHandler := topuri.NewTopUriMetricHandler(crCache)
	http.Handle(metricUrl, topUriMetricHandler)

	server := http.Server{
		Addr:              *ListenAddress,
	}
	if err := server.ListenAndServe();err != nil{
		logrus.Panic(err)
	}

}
