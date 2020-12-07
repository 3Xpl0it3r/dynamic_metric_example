package topuri

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"nginxexporter/pkg/collector"
	"nginxexporter/pkg/crcache"
)

type handler struct {
	// cache
	cache *crcache.CrCache
	enableInternalMetric bool
}



func NewTopUriMetricHandler(cache *crcache.CrCache)*handler{
	return &handler{
		cache: cache,
		enableInternalMetric: false,
	}
}

func(h *handler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	register := prometheus.NewRegistry()
	dynamicMetrics := generateAllMetricItems(h.cache)
	registerMetric(register, dynamicMetrics)

	handler := promhttp.HandlerFor(register, promhttp.HandlerOpts{
		Registry:            register,
		EnableOpenMetrics:   false,
	})
	handler.ServeHTTP(w, r)
}


func generateAllMetricItems(cache *crcache.CrCache)[]prometheus.Counter{
	var metricItems = []prometheus.Counter{}
	for item, _ := range cache.Data{
		uriMetric := prometheus.NewCounter(prometheus.CounterOpts{Namespace: collector.NginxCollectorExporterNamespace,
			Subsystem: SubSystem, Name: item, Help: fmt.Sprintf("the number of %s visited", item)})
		metricItems = append(metricItems, uriMetric)
	}
	return metricItems
}

func registerMetric(register prometheus.Registerer,metricItems []prometheus.Counter){
	for _, metric := range metricItems{
		metric.Add(10)
		register.MustRegister(metric)
	}
}