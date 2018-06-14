package servicea

import (
	"context"
	"log"

	api "github.com/shihanng/gaegoasample/svc/servicea/gen/api"
)

// api service example implementation.
// The example methods log the requests and return zero values.
type apiSvc struct {
	logger *log.Logger
}

// NewAPI returns the api service implementation.
func NewAPI(logger *log.Logger) api.Service {
	return &apiSvc{logger}
}

// Show info of the service
func (s *apiSvc) Info(ctx context.Context) (res *api.GaegoasampleInfo, err error) {
	res = &api.GaegoasampleInfo{}
	s.logger.Print("api.info")
	return
}
