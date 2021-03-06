// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// api service
//
// Command:
// $ goa gen github.com/shihanng/gaegoasample/svc/servicea/design -o
// /home/shihanng/go/src/github.com/shihanng/gaegoasample/svc/servicea

package api

import (
	"context"
)

// API of the service
type Service interface {
	// Show info of the service
	Info(context.Context) (res *GaegoasampleInfo, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "api"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"info"}

// GaegoasampleInfo is the result type of the api service info method.
type GaegoasampleInfo struct {
	// ID of the service
	ID *string
	// Service's name
	ServiceName *string
	// Service's version
	Version *string
}
