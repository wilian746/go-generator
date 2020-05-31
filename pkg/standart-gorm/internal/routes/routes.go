package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
	ServerConfig "github.com/wilian746/go-generator/pkg/standart-gorm/configs"
	HealthHandler "github.com/wilian746/go-generator/pkg/standart-gorm/internal/handlers/health"
	ProductHandler "github.com/wilian746/go-generator/pkg/standart-gorm/internal/handlers/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
)

const BasePath = "/api/v1"

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters()

	r.RouterSwagger()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) setConfigsRouters() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterSwagger() {
	swaggerHost := fmt.Sprintf("http://localhost:%v/swagger/doc.json", ServerConfig.GetConfig().Port)
	r.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerHost),
	))
}

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.NewHandler(repository)

	r.router.Route(BasePath+"/health", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.NewHandler(repository)

	r.router.Route(BasePath+"/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Get("/{ID}", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
