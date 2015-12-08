package stack

import "net/http"

type chainHandler func(*Context) http.Handler
type chainMiddleware func(*Context, http.Handler) http.Handler

type Chain struct {
	mws []chainMiddleware
	h   chainHandler
}

func New(mws ...chainMiddleware) Chain {
	return Chain{mws: mws}
}

func (c Chain) Append(mws ...chainMiddleware) Chain {
	newMws := make([]chainMiddleware, len(c.mws)+len(mws))
	copy(newMws[:len(c.mws)], c.mws)
	copy(newMws[len(c.mws):], mws)
	c.mws = newMws
	return c
}

type HandlerChain struct {
	context *Context
	Chain
}

func adaptHandler(h http.Handler) chainHandler {
	return func(ctx *Context) {
		return h
	}
}

func adaptHandlerFunc(fn func(w http.ResponseWriter, r *http.Request)) chainHandler {
	return adaptHandler(http.HandlerFunc(fn))
}

func adaptContextHandlerFunc(fn func(ctx *Context, w http.ResponseWriter, r *http.Request)) chainHandler {
	return func(ctx *Context) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn(ctx, w, r)
		})
	}
}

func (c Chain) Then(chf func(ctx *Context, w http.ResponseWriter, r *http.Request)) HandlerChain {
	c.h = adaptContextHandlerFunc(chf)
	return newHandlerChain(c)
}
