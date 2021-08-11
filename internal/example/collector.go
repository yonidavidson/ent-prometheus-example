package example

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// List of dynamic labels
var labelNames = []string{"endpoint", "error_code"}

// Create a counter collector
var exampleCollector = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name:        "endpoint_errors",
		Help:        "Number of errors in endpoints",
	},
	labelNames,
)

// When using you set the values of the dynamic labels and then increment the counter
func incrementError(){
	exampleCollector.WithLabelValues("/create-user","400").Inc()
}