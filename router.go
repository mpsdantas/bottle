package bottle

import (
	"github.com/gofiber/fiber/v2"
)

type Handler = func(c *HandlerCtx) error

type Router struct {
	router fiber.Router
}

func newRouter(router fiber.Router) *Router {
	return &Router{router: router}
}

func (r *Router) Use(args ...interface{}) {
	var handlers []Handler

	for _, arg := range args {
		if h, ok := arg.(Handler); ok {
			handlers = append(handlers, h)
		}
	}

	r.router.Use(toEngineHandlersInterface(handlers)...)
}

func (r *Router) Get(path string, handlers ...Handler) {
	r.router.Get(path, toEngineHandlers(handlers)...)
}

func (r *Router) Post(path string, handlers ...Handler) {
	r.router.Post(path, toEngineHandlers(handlers)...)
}

func (r *Router) Put(path string, handlers ...Handler) {
	r.router.Put(path, toEngineHandlers(handlers)...)
}

func (r *Router) Delete(path string, handlers ...Handler) {
	r.router.Delete(path, toEngineHandlers(handlers)...)
}

func (r *Router) Patch(path string, handlers ...Handler) {
	r.router.Patch(path, toEngineHandlers(handlers)...)
}

func (r *Router) Options(path string, handlers ...Handler) {
	r.router.Options(path, toEngineHandlers(handlers)...)
}

func (r *Router) Head(path string, handlers ...Handler) {
	r.router.Head(path, toEngineHandlers(handlers)...)
}

func (r *Router) Connect(path string, handlers ...Handler) {
	r.router.Connect(path, toEngineHandlers(handlers)...)
}

func (r *Router) Trace(path string, handlers ...Handler) {
	r.router.Trace(path, toEngineHandlers(handlers)...)
}

func (r *Router) Group(prefix string, handlers ...Handler) *Router {
	group := r.router.Group(prefix, toEngineHandlers(handlers)...)
	return &Router{router: group}
}

func toEngineHandlers(handlers []Handler) []fiber.Handler {
	var res []fiber.Handler

	for _, handler := range handlers {
		h := handler
		res = append(res, func(ctx *fiber.Ctx) error {
			return h(WrapEngineContext(ctx))
		})
	}

	return res
}

func toEngineHandlersInterface(handlers []Handler) []interface{} {
	var res []interface{}

	for _, handler := range handlers {
		h := handler
		res = append(res, func(ctx *fiber.Ctx) error {
			return h(WrapEngineContext(ctx))
		})
	}

	return res
}
