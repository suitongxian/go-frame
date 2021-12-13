package engine

import (
	"net/http"
	"strings"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	*GroupRouter
	Router *Router
	groups []*GroupRouter
}

func New() *Engine {
	engine := &Engine{Router: &Router{Handler: make(map[string]HandlerFunc)}}
	engine.GroupRouter = &GroupRouter{Eng: engine}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares = make([]HandlerFunc, 0, len(e.Middlewares))
	for i := 0;i < len(e.groups);i++ {
		if strings.HasPrefix(req.URL.Path, e.groups[i].Prefix) {
			middlewares = append(middlewares, e.groups[i].Middlewares...)
		}
	}
	ctx := NewContext(w, req)
	ctx.Handlers = middlewares
	e.Router.handle(ctx)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
