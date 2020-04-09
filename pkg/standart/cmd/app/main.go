package main

import (
	"fmt"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/instance"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/config"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/routes"
	"log"
	"net/http"
)

func main() {
	configs := config.GetConfig()
	entity := &product.Product{}

	connection := instance.GetConnection(configs.Dialect, configs.DatabaseURI)
	connection.Table(entity.TableName()).AutoMigrate(entity)
	repository := adapter.NewAdapter(connection)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	log.Println("service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}
