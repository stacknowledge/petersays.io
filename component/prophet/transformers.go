package prophet

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeProphesizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type EnlightmentRequest struct {
	Saying string `json:"saying" binding:"required"`
}

func DecodeEnlightmentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EnlightmentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

type ProphesizeResponse struct {
	Saying string `json:"saying"`
}

type EnlightmentResponse struct {
	Status  string `json:"status,omitempty"`
	Saying  string `json:"saying"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func EncodeProphetResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(response)
}

type ProphetErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func EncodeProphetErrorResponse(context context.Context, err error, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(&ProphetErrorResponse{err.Error(), "An wheel splited from the engine. A monkey will fix the problem asap!"})
}
