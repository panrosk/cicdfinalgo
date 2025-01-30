package main

import (
	"cicdfinalgo/metrics"
	"cicdfinalgo/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"runtime"
	"time"
)

func main() {
	metrics.RegisterMetrics()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		metrics.ActiveConnections.Inc()
		defer metrics.ActiveConnections.Dec()

		who := utils.SayHiTo("Raymundo y todo el mundo")
		fmt.Fprint(w, who)

		metrics.HttpRequestsTotal.WithLabelValues("/hello").Inc()

		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("/hello").Observe(duration)

		responseSize := len(who)
		metrics.HttpResponseSizeBytes.WithLabelValues("/hello").Observe(float64(responseSize))
	})

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			metrics.MemoryUsage.Set(float64(memStats.Alloc))
			metrics.CPUUsage.Observe(float64(runtime.NumGoroutine()))
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
