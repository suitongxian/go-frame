package engine

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	*GroupRouter
	Router *Router
}

func New() *Engine {
	engine := &Engine{Router: &Router{Handler: make(map[string]HandlerFunc)}}
	engine.GroupRouter = &GroupRouter{Eng: engine}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	e.Router.handle(ctx)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
