package entprom

import (
	"context"
	"time"

	"entgo.io/ent"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//Ent dynamic dimensions
const (
	mutationType = "mutation_type"
	mutationOp   = "mutation_op"
)

var entLabels = []string{mutationType, mutationOp}

// Create a collector for total operations counter
func initOpsProcessedTotal() *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ent_operation_total",
			Help: "Number of ent mutation operations",
		},
		entLabels,
	)
}

// Create a collector for error counter
func initOpsProcessedError() *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ent_operation_error",
			Help: "Number of failed ent mutation operations",
		},
		entLabels,
	)
}

// Create a collector for duration histogram collector
func initOpsDuration() *prometheus.HistogramVec {
	return promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ent_operation_duration_seconds",
			Help: "Time in seconds per operation",
		},
		entLabels,
	)
}

// Hook init collectors, count total at beginning error on mutation error and duration also after.
func Hook() ent.Hook {
	opsProcessedTotal := initOpsProcessedTotal()
	opsProcessedError := initOpsProcessedError()
	opsDuration := initOpsDuration()
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Before mutation, start measuring time.
			start := time.Now()
			// Extract dynamic labels from mutation.
			labels := prometheus.Labels{mutationType: m.Type(), mutationOp: m.Op().String()}
			// Increment total ops counter.
			opsProcessedTotal.With(labels).Inc()
			// Execute mutation.
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// In case of error increment error counter.
				opsProcessedError.With(labels).Inc()
			}
			// Stop time measure.
			duration := time.Since(start)
			// Record duration in seconds.
			opsDuration.With(labels).Observe(duration.Seconds())
			return v, err
		})
	}
}
