package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"log"
)

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "rpc_durations_ms",
			Help: "RPC latency distributions.",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			},
		},
	)
	rpcRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "rpc_requests",
		Help: "Service requests.",
	})
)

func init() {
	// Register the summary and counter with Prometheus's default registry.
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcRequests)



	go func() {
		log.Printf("Starting metrics exporter at http://localhost%v.\n", ":9102")
		log.Fatal(http.ListenAndServe(":9102", promhttp.Handler()))
	}()
}
