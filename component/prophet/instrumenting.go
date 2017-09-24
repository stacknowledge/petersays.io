package prophet

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
			Namespace: "petersays",
			Subsystem: "prophet",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "petersays",
			Subsystem: "prophet",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		service,
	}
}

func (middleware InstrumentingMiddleware) Prophesize() string {
	defer func(begin time.Time) {
		lvs := []string{"method", "Prophet::prophesize", "error", ""}
		middleware.requestCount.With(lvs...).Add(1)
		middleware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middleware.next.Prophesize()
}

func (middleware InstrumentingMiddleware) Enlightment(enlightment string) (string, error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Prophet::enlightment", "error", ""}
		middleware.requestCount.With(lvs...).Add(1)
		middleware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middleware.next.Enlightment(enlightment)
}
