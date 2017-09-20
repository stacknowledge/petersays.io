package instrumenting

import (
	"microservice/services"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
)

type PetersaysInstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           services.PetersaysService
}

func NewPetersaysInstrumentingMiddleware(next services.PetersaysService) *PetersaysInstrumentingMiddleware {
	return &PetersaysInstrumentingMiddleware{
		RequestCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "Petersays",
			Subsystem: "petersays_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method", "error"}),

		RequestLatency: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "Petersays",
			Subsystem: "petersays_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method", "error"}),

		CountResult: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "Petersays",
			Subsystem: "petersays_service",
			Name:      "count_result",
			Help:      "The result of each count method.",
		}, []string{}),
	}
}
