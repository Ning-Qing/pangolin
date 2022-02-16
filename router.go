package pangolin

// Handler
// 请求处理
type Handler func(ctx *PangolinCtx)

type Router struct {
	router map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		router: make(map[string]Handler),
	}
}

func (r *Router) Add(method string, handler Handler) {
	r.router[method] = handler
}

func (r *Router)run(method string)Handler{
	return r.router[method]
}