package handlers

import (
	"expvar"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/duexcoast/webservice/app/services/sales-api/handlers/debug/checkgrp"
	"github.com/duexcoast/webservice/app/services/sales-api/handlers/v1/testgrp"
	"github.com/duexcoast/webservice/business/web/mid"
	"github.com/duexcoast/webservice/foundation/web"
	"go.uber.org/zap"
)

// StandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard libary debug endpoints
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

// DebugMux registers all the debug standard library routes and the custom
// debug application routes for the service, thus bypassing the use of the
// DefaultServerMux. Using the DefaultServerMux would be a security risk since
// a dependendency could inject a handler into our service without us knowing it.
func DebugMux(build string, log *zap.SugaredLogger) http.Handler {
	mux := DebugStandardLibraryMux()

	// Register debug check endpoints.
	cgh := checkgrp.Handlers{
		Build: build,
		Log:   log,
	}
	mux.HandleFunc("/debug/readiness", cgh.Readiness)
	mux.HandleFunc("/debug/liveness", cgh.Liveness)

	return mux
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
	)

	// Load the routes for teh different versions of the API.
	v1(app, cfg)

	return app
}

// v1 binds all the version 1 routes
func v1(app *web.App, cfg APIMuxConfig) {
	const version = "v1"
	tgh := testgrp.Handlers{
		Log: cfg.Log,
	}

	app.Handle(http.MethodGet, "v1", "/test", tgh.Test)

}
