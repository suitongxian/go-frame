package engine

import "fmt"

type Router struct {
	Handler map[string]HandlerFunc
}

func (r *Router) AddRouter(method string, url string, handler HandlerFunc) {
	var key string
	key = CreateRouterKey(method, url)
	r.Handler[key] = handler
}

func (r *Router) handle(ctx *Context)  {
	var key string
	key = CreateRouterKey(ctx.Method, ctx.Path)
	if v, ok := r.Handler[key]; ok {
		ctx.Handlers = append(ctx.Handlers, v)
	} else {
		fmt.Printf("404 NOT Found: %s", ctx.Path)
	}
	ctx.Next()
}