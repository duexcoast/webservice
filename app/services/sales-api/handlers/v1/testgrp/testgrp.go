// Package testgrp contains all the test handlers.
package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/duexcoast/webservice/business/sys/validate"
	"github.com/duexcoast/webservice/foundation/web"
	"go.uber.org/zap"
)

// Handlers mangaes the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		// panic("testing panic")
		return validate.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		// return errors.New("untrusted error")
	}

	status := struct {
		Status string
	}{
		Status: "HELL YEAH",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
