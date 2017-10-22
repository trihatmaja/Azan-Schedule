package handler

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	once   sync.Once
	histo1 *prometheus.HistogramVec
	histo2 *prometheus.HistogramVec
)

func init() {
	once.Do(func() {
		histo1 = newHistogramVec()
		histo2 = newHistogramVecCache()
		prometheus.MustRegister(histo1)
		prometheus.MustRegister(histo2)
	})
}

// Handler is an http.HandlerFunc to serve metrics endpoint
func PromHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

// TraceRequestTime is method to write metrics to prometheus
func TraceRequestTime(method, action, status string, elapsedTime float64) {
	histo1.WithLabelValues(method, action, status).Observe(elapsedTime)
}

func TraceRequestTimeCache(method, action, cache string, elapsedTime float64) {
	histo2.WithLabelValues(method, action, cache).Observe(elapsedTime)
}

func newHistogramVec() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "service_latency_seconds",
		Help: "service events response in miliseconds",
	}, []string{"method", "action", "status"})
}

func newHistogramVecCache() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "service_latency_seconds_cache",
		Help: "service events response in miliseconds with cache",
	}, []string{"method", "action", "cache"})
}
