package prophet

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

func makeProphesizeEndpoint(service Service, logger *logrus.Logger) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {

		saying := service.Prophesize()

		return ProphesizeResponse{
			saying,
		}, nil
	}
}

func makeEnlightmentEndpoint(service Service, logger *logrus.Logger) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		enlightmentRequest := request.(EnlightmentRequest)

		if enlightmentRequest.Saying == "" {
			return EnlightmentResponse{
				Status: "rejected",
				Error:  "Peters does not accept those kinds of enlightments",
			}, nil
		}

		saying, error := service.Enlightment(enlightmentRequest.Saying)

		if error != nil {
			logger.WithFields(logrus.Fields{
				"Method":    "Prophet::enlightment",
				"Parameter": enlightmentRequest.Saying,
			}).Error(error.Error())

			return EnlightmentResponse{
				Saying: saying,
				Status: "unaccepted",
				Error:  error.Error(),
			}, nil
		}

		return EnlightmentResponse{
			Saying:  saying,
			Status:  "Pending",
			Message: "Your enlightment will be considered. Thank you, my kid!",
		}, nil
	}
}
