package example

import (
	"net/http"

	"github.com/sirupsen/logrus"

	httptransport "github.com/go-kit/kit/transport/http"
)

type ExampleWheel struct {
	domain  string
	service Service
}

func NewWheel(logger *logrus.Logger) *ExampleWheel {
	var service Service

	service = &ExampleService{}
	service = NewLoggingMiddleware(logger, service)
	service = NewInstrumentingMiddleware(service)

	wheel := &ExampleWheel{
		"/example",
		service,
	}

	return wheel
}

func (wheel *ExampleWheel) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle(wheel.domain+"/", httptransport.NewServer(
		makeMethodEndpoint(wheel.service),
		DecodeMethodResponse,
		EncodeMethodResponse,
	))
}
