package endpoints

import "github.com/go-kit/kit/endpoint"

type PetersaysEndpoints struct {
	Say endpoint.Endpoint
}

func (endpoints *PetersaysEndpoints) PeterSay() string {
	return ""
}
