package engine

type GroupRouter struct {
	Prefix string
	Eng    *Engine
	//parent *GroupRouter
}

func (g *GroupRouter) Group(prefix string) *GroupRouter {
	e := g.Eng
	return &GroupRouter{
		Prefix: prefix,
		Eng: e,
	}
}

func (g *GroupRouter) Get(url string, handler HandlerFunc) {
	g.Eng.Router.AddRouter("GET", url, handler)
}

func (g *GroupRouter) Post(url string, handler HandlerFunc) {
	g.Eng.Router.AddRouter("POST", url, handler)
}