package main

import (
	"fmt"
	"github.com/wilian746/go-generator/pkg/standart-gorm/config"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/routes"
	"github.com/wilian746/go-generator/pkg/standart-gorm/repository/adapter"
	"github.com/wilian746/go-generator/pkg/standart-gorm/repository/database"
	"log"
	"net/http"
)

func main() {
	configs := config.GetConfig()
	entity := &product.Product{}

	connection := database.GetConnection(configs.Dialect, configs.DatabaseURI)
	connection.Table(entity.TableName()).AutoMigrate(entity)
	repository := adapter.NewAdapter(connection)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	log.Println("service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}
