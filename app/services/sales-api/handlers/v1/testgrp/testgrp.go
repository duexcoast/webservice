package testgrp

import (
	"context"
	"net/http"

	"github.com/duexcoast/webservice/foundation/web"
	"go.uber.org/zap"
)

// Handlers mangaes the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "HELL YEAH",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
