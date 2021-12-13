package engine

type GroupRouter struct {
	Prefix string
	Eng    *Engine
	Middlewares []HandlerFunc
	//parent *GroupRouter
}

func (g *GroupRouter) Group(prefix string) *GroupRouter {
	e := g.Eng
	group := &GroupRouter{
		Prefix: g.Prefix + prefix,
		Eng: e,
	}
	e.groups = append(e.groups, group)
	return group
}

func (g *GroupRouter) Use(middlewares ...HandlerFunc) {
	g.Middlewares = append(g.Middlewares, middlewares...)
}

func (g *GroupRouter) Get(url string, handler HandlerFunc) {
	url = SplitUrl(g.Prefix, url)
	g.Eng.Router.AddRouter("GET", url, handler)
}

func (g *GroupRouter) Post(url string, handler HandlerFunc) {
	url = SplitUrl(g.Prefix, url)
	g.Eng.Router.AddRouter("POST", url, handler)
}