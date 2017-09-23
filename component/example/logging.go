package example

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

func (middleware LoggingMiddleware) Method() string {

	defer func(begin time.Time) {
		middleware.logger.WithFields(logrus.Fields{
			"Method": "ExampleService::method",
			"Took":   time.Since(begin),
		}).Debug("Logging report")
	}(time.Now())

	return middleware.next.Method()
}
