package logging

import (
	"microservice/services"
	"time"

	"github.com/sirupsen/logrus"
)

type PetersaysLoggingMiddleware struct {
	logger *logrus.Logger
	next   services.PetersaysService
}

func (middleware *PetersaysLoggingMiddleware) PeterSay() string {

	defer func(begin time.Time) {
		middleware.logger.WithFields(logrus.Fields{
			"method": "Petersays.PeterSay",
			"took":   time.Since(begin),
		}).Info("An request was received to petersays service.")

	}(time.Now())

	return middleware.next.PeterSay()
}
