package example

import (
	"time"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

func NewInstrumentingMiddleware(service Service) *InstrumentingMiddleware {
	fieldKeys := []string{"method", "error"}

	return &InstrumentingMiddleware{
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "example",
			Subsystem: "example_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "example",
			Subsystem: "example_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		service,
	}
}

func (middleware InstrumentingMiddleware) Method() string {
	defer func(begin time.Time) {
		lvs := []string{"method", "ExampleService::method", "error", ""}
		middleware.requestCount.With(lvs...).Add(1)
		middleware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middleware.next.Method()
}
