// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// api HTTP server types
//
// Command:
// $ goa gen github.com/shihanng/gaegoasample/svc/servicea/design -o
// /home/shihanng/go/src/github.com/shihanng/gaegoasample/svc/servicea

package server

import (
	api "github.com/shihanng/gaegoasample/svc/servicea/gen/api"
)

// InfoResponseBody is the type of the "api" service "info" endpoint HTTP
// response body.
type InfoResponseBody struct {
	// ID of the service
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Service's name
	ServiceName *string `form:"service_name,omitempty" json:"service_name,omitempty" xml:"service_name,omitempty"`
	// Service's version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// NewInfoResponseBody builds the HTTP response body from the result of the
// "info" endpoint of the "api" service.
func NewInfoResponseBody(res *api.GaegoasampleInfo) *InfoResponseBody {
	body := &InfoResponseBody{
		ID:          res.ID,
		ServiceName: res.ServiceName,
		Version:     res.Version,
	}
	return body
}
