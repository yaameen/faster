package faster

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Ctx = fiber.Ctx
type Handler = fiber.Handler
type Static = fiber.Static
type Config = fiber.Config
type TLSHandler = fiber.TLSHandler

type FastRouter interface {
	Static(prefix string, root string, config ...Static) FastRouter
	Add(method, path string, handlers ...Handler) FastRouter
	Options(path string, handlers ...Handler) FastRouter
	Trace(path string, handlers ...Handler) FastRouter
	Get(path string, handlers ...Handler) FastRouter
	Post(path string, handlers ...Handler) FastRouter
	Patch(path string, handlers ...Handler) FastRouter
	Put(path string, handlers ...Handler) FastRouter
	Head(path string, handlers ...Handler) FastRouter
	Delete(path string, handlers ...Handler) FastRouter
	Connect(path string, handlers ...Handler) FastRouter
	Use(args ...interface{}) FastRouter
	All(path string, handlers ...Handler) FastRouter
	Any(path string, handlers ...Handler) FastRouter
	Group(handlers ...Handler) FastRouter
	Prefix(prefix string, handlers ...Handler) FastRouter
	Mount(prefix string, router FastRouter) FastRouter
}

type FastApp struct {
	app *fiber.App
}

func (h *FastApp) Test(req *http.Request, msTimeout ...int) (resp *http.Response, err error) {
	return h.app.Test(req, msTimeout...)
}

func (h *FastApp) Static(prefix string, root string, config ...Static) FastRouter {
	h.app.Static(prefix, root, config...)
	return h
}

func (h *FastApp) Add(method, path string, handlers ...Handler) FastRouter {
	h.app.Add(method, path, handlers...)
	return h
}

func (h *FastApp) Options(path string, handlers ...Handler) FastRouter {
	h.app.Options(path, handlers...)
	return h
}

func (h *FastApp) Trace(path string, handlers ...Handler) FastRouter {
	h.app.Trace(path, handlers...)
	return h
}

func (h *FastApp) Get(path string, handlers ...Handler) FastRouter {
	h.app.Get(path, handlers...)
	return h
}

func (h *FastApp) Post(path string, handlers ...Handler) FastRouter {
	h.app.Post(path, handlers...)
	return h
}

func (h *FastApp) Patch(path string, handlers ...Handler) FastRouter {
	h.app.Patch(path, handlers...)
	return h
}

func (h *FastApp) Put(path string, handlers ...Handler) FastRouter {
	h.app.Put(path, handlers...)
	return h
}

func (h *FastApp) Head(path string, handlers ...Handler) FastRouter {
	h.app.Head(path, handlers...)
	return h
}

func (h *FastApp) Delete(path string, handlers ...Handler) FastRouter {
	h.app.Delete(path, handlers...)
	return h
}

func (h *FastApp) Connect(path string, handlers ...Handler) FastRouter {
	h.app.Connect(path, handlers...)
	return h
}

func (h *FastApp) Use(args ...interface{}) FastRouter {
	h.app.Use(args...)
	return h
}

func (h *FastApp) All(path string, handlers ...Handler) FastRouter {
	h.app.All(path, handlers...)
	return h
}

func (h *FastApp) Any(path string, handlers ...Handler) FastRouter {
	return h.All(path, handlers...)
}

func (h *FastApp) Mount(prefix string, router FastRouter) FastRouter {
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}
	if app, ok := router.(*FastApp); ok {
		println(fmt.Sprintf("mount %s", prefix))
		h.app.Mount(prefix, app.app)
	}
	return h.Prefix(prefix)
}

func (h *FastApp) Group(handlers ...Handler) FastRouter {
	grp := &FastGroup{
		handlers: handlers,
		router:   h,
	}
	return grp
}

func (h *FastApp) Prefix(prefix string, handlers ...Handler) FastRouter {
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}
	grp := &FastGroup{
		handlers: handlers,
		router:   h,
		prefix:   &prefix,
	}
	return grp
}

func (r *FastApp) SetTLSHandler(handler TLSHandler) {
	r.app.SetTLSHandler(&handler)
}

func (r *FastApp) Config() Config {
	return r.app.Config()
}

func (r *FastApp) Handler() fasthttp.RequestHandler {
	return r.app.Handler()
}

func (r *FastApp) Stack() [][]*fiber.Route {
	return r.app.Stack()
}

func (r *FastApp) HandlersCount() uint32 {
	return r.app.HandlersCount()
}

func (r *FastApp) Shutdown() error {
	return r.app.Shutdown()
}

func (r *FastApp) Server() *fasthttp.Server {
	return r.app.Server()
}

func (r *FastApp) Hooks() *fiber.Hooks {
	return r.app.Hooks()
}

func (r *FastApp) ErrorHandler(ctx *fiber.Ctx, err error) error {
	return r.app.ErrorHandler(ctx, err)
}

func (r *FastApp) Listen(addr string) error {
	return r.app.Listen(addr)
}

type FastGroup struct {
	handlers []Handler
	router   *FastApp
	prefix   *string
}

func (f FastGroup) fixPrefix(path string) string {
	if f.prefix != nil {
		if path[0] == '/' {
			return *f.prefix + path
		}
		return *f.prefix + "/" + path
	}
	return path
}

func (h *FastGroup) Static(prefix string, root string, config ...Static) FastRouter {
	h.router.Static(h.fixPrefix(prefix), root, config...)
	return h
}

func (h *FastGroup) Add(method, path string, handlers ...Handler) FastRouter {
	h.router.Add(method, h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Options(path string, handlers ...Handler) FastRouter {
	h.router.Options(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Trace(path string, handlers ...Handler) FastRouter {
	h.router.app.Trace(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Get(path string, handlers ...Handler) FastRouter {
	h.router.Get(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Post(path string, handlers ...Handler) FastRouter {
	h.router.Post(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Patch(path string, handlers ...Handler) FastRouter {
	h.router.Patch(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Put(path string, handlers ...Handler) FastRouter {
	h.router.Put(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Head(path string, handlers ...Handler) FastRouter {
	h.router.Head(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Delete(path string, handlers ...Handler) FastRouter {
	h.router.Delete(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Connect(path string, handlers ...Handler) FastRouter {
	h.router.Connect(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Use(args ...interface{}) FastRouter {
	h.router.Use(args...)
	return h
}

func (h *FastGroup) All(path string, handlers ...Handler) FastRouter {
	h.router.All(h.fixPrefix(path), append(h.handlers, handlers...)...)
	return h
}

func (h *FastGroup) Any(path string, handlers ...Handler) FastRouter {
	return h.All(path, handlers...)
}

func New(config ...Config) *FastApp {
	return &FastApp{
		app: fiber.New(config...),
	}
}

func (h *FastGroup) Mount(prefix string, router FastRouter) FastRouter {
	p := h.fixPrefix(prefix)
	group := &FastGroup{prefix: &p, router: h.router}
	for _, v := range router.(*FastApp).app.Stack() {
		for _, r := range v {
			group.Add(r.Method, r.Path, append(h.handlers, r.Handlers...)...)
		}
	}
	return group
}

func (r *FastGroup) Group(handlers ...Handler) FastRouter {
	grp := &FastGroup{
		handlers: append(r.handlers, handlers...),
		router:   r.router,
		prefix:   r.prefix,
	}
	return grp
}

func (r *FastGroup) Prefix(prefix string, handlers ...Handler) FastRouter {
	prefix = r.fixPrefix(prefix)
	grp := &FastGroup{
		handlers: append(r.handlers, handlers...),
		router:   r.router,
		prefix:   &prefix,
	}
	return grp
}
