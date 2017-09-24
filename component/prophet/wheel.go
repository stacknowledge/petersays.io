package prophet

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	httptransport "github.com/go-kit/kit/transport/http"
)

type ProphetWheel struct {
	domain  string
	service Service
	logger  *logrus.Logger
}

func NewWheel(logger *logrus.Logger) *ProphetWheel {
	var service Service

	service = &ProphetService{}
	service = NewLoggingMiddleware(logger, service)
	service = NewInstrumentingMiddleware(service)

	wheel := &ProphetWheel{
		"/api/peters",
		service,
		logger,
	}

	return wheel
}
func (wheel *ProphetWheel) RegisterHandlers(mux *mux.Router) {
	mux.Handle(wheel.domain+"/saying", httptransport.NewServer(
		makeProphesizeEndpoint(wheel.service, wheel.logger),
		DecodeProphesizeRequest,
		EncodeProphetResponse,
		httptransport.ServerErrorEncoder(EncodeProphetErrorResponse),
	)).Methods("GET")

	mux.Handle(wheel.domain+"/enlightment", httptransport.NewServer(
		makeEnlightmentEndpoint(wheel.service, wheel.logger),
		DecodeEnlightmentRequest,
		EncodeProphetResponse,
		httptransport.ServerErrorEncoder(EncodeProphetErrorResponse),
	)).Methods("POST")
}
