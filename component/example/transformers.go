package example

import (
	"context"
	"encoding/json"
	"net/http"
)

type MethodResponse struct {
	Example string `json:"example"`
}

func DecodeMethodResponse(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeMethodResponse(_ context.Context, w http.ResponseWriter, methodResponse interface{}) error {
	return json.NewEncoder(w).Encode(methodResponse)
}
