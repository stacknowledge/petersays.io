package transformers

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	PeterSayResponse struct {
		Say   string `json:"say"`
		Error string `json:"error,omitempty"`
	}
)

func encodePetersaysResponse(context context.Context, writter http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writter).Encode(response)
}
