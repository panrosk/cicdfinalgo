package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de solicitudes HTTP recibidas",
		},
		[]string{"path"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duración de las solicitudes HTTP en segundos",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	HttpRequestErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total de errores HTTP",
		},
		[]string{"path", "status_code"},
	)

	ActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Número actual de conexiones activas",
		},
	)

	HttpResponseSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_response_size_bytes",
			Help:       "Tamaño de las respuestas HTTP en bytes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path"},
	)

	CPUUsage = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "cpu_usage_seconds",
			Help:    "Uso de CPU en segundos",
			Buckets: prometheus.DefBuckets,
		},
	)

	MemoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Memoria utilizada por la aplicación en bytes",
		},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(HttpRequestErrorsTotal)
	prometheus.MustRegister(ActiveConnections)
	prometheus.MustRegister(HttpResponseSizeBytes)
	prometheus.MustRegister(CPUUsage)
	prometheus.MustRegister(MemoryUsage)
}
