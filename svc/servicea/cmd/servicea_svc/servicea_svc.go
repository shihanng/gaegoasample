package servicea_svc

import (
	"context"
	"log"
	"net/http"
	"os"

	servicea "github.com/shihanng/gaegoasample/svc/servicea"
	api "github.com/shihanng/gaegoasample/svc/servicea/gen/api"
	apisvr "github.com/shihanng/gaegoasample/svc/servicea/gen/http/api/server"
	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
	"google.golang.org/appengine"
)

func init() {
	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		adapter middleware.Logger
		logger  *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[servicea] ", log.Ltime)
		adapter = middleware.NewLogger(logger)
	}

	// Create the structs that implement the services.
	var (
		apiSvc api.Service
	)
	{
		apiSvc = servicea.NewAPI(logger)
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		apiEndpoints *api.Endpoints
	)
	{
		apiEndpoints = api.NewEndpoints(apiSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		apiServer *apisvr.Server
	)
	{
		eh := ErrorHandler(logger)
		apiServer = apisvr.New(apiEndpoints, mux, dec, enc, eh)
	}

	// Configure the mux.
	apisvr.Mount(mux, apiServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = middleware.Log(adapter)(handler)
		handler = middleware.RequestID()(handler)
		handler = appEngineContext()(handler)
	}

	http.Handle("/", handler)
}

func appEngineContext() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := appengine.NewContext(r)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
