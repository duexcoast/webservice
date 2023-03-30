package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/duexcoast/webservice/foundation/web"
	"go.uber.org/zap"
)

// Logger...
func Logger(log *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, err := web.GetValues(ctx)
			if err != nil {
				return err // web.NewShutdownError("web value missing from context")
			}

			log.Infow("request started", "traceid", v.traceID, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err = handler(ctx, w, r)

			log.Infow("request completed", "traceid", v.traceID, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr, "statuscode", v.statuscode, "since", time.Since(now))

			return err
		}
		return h
	}

	return m
}
