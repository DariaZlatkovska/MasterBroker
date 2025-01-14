package transport

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var request = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Total number of http request",
	},
	[]string{"method", "endpoint"},
)

func InitPrometheus() {
	prometheus.MustRegister(request)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request.WithLabelValues(r.Method, r.URL.Path).Inc()
		next.ServeHTTP(w, r)
	})
}
