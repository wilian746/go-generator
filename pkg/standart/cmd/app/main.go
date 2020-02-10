package main

import (
	"github.com/go-chi/chi"
	"github.com/wilian746/gorm-crud-generator/internal/utils/env"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/instance"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	ProductHandler "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/handlers/product"
	"log"
	"net/http"
)

func main() {
	port := env.GetEnv("PORT", "8666")
	entity := &product.Product{}

	connection := instance.GetConnection("sqlite3", ":memory:")

	connection.Table(entity.TableName()).AutoMigrate(entity)

	repository := adapter.NewAdapter(connection)

	log.Println("service running on port :", port)

	log.Fatal(http.ListenAndServe(":" + port, Router(repository)))
}

func Router(repository adapter.Interface) *chi.Mux {
	router := chi.NewRouter()

	RouterProduct(router, repository)

	return router
}

func RouterProduct (router *chi.Mux, repository adapter.Interface) *chi.Mux {
	handler := ProductHandler.NewHandler(repository)

	router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Get("/{ID}", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})

	return router
}