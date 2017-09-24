package prophet

import (
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	logger *logrus.Logger
	next   Service
}

func NewLoggingMiddleware(logger *logrus.Logger, next Service) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger,
		next,
	}
}

func (middleware LoggingMiddleware) Prophesize() string {
	defer func(begin time.Time) {
		middleware.logger.WithFields(logrus.Fields{
			"Method": "Prophet::prophesize",
			"Took":   time.Since(begin),
		}).Debug("Logging report")
	}(time.Now())

	return middleware.next.Prophesize()
}

func (middleware LoggingMiddleware) Enlightment(enlightment string) (string, error) {
	defer func(begin time.Time) {
		middleware.logger.WithFields(logrus.Fields{
			"Method": "Prophet::enlightment",
			"Took":   time.Since(begin),
		}).Debug("Logging report")
	}(time.Now())

	return middleware.next.Enlightment(enlightment)
}
