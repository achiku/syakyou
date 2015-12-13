package xhandler

import (
	"net/http"

	"golang.org/x/net/context"
)

// HandlerC is net/context aware http.Handler
type HandlerC interface {
	ServeHTTPC(context.Context, http.ResponseWriter, *http.Request)
}

// HandlerFuncC type is an adapter to allow the use of ordinaly functions
// as a xhandler.Handler. If f is a function with appropriate signature,
// xhandler.HandlerFunc(f is a xhandler.Handler object that calls f.
type HandlerFuncC func(context.Context, http.ResponseWriter, *http.Request)

// ServeHTTPC calls f(ctx, w, r)
func (f HandlerFuncC) ServeHTTPC(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}

// New creates a conventional http.Handler injecting the provided root
// context to sub handlers. This handler is used as a bridge between conventional
// http.Handler and context awarte handlers.
func New(ctx context.Context, h HandlerC) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTPC(ctx, w, r)
	})
}
