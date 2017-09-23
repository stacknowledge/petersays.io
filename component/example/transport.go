package example

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeMethodEndpoint(service Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {

		stringExample := service.Method()

		return MethodResponse{
			stringExample,
		}, nil
	}
}
