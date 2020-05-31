package main

import (
	"fmt"
	"github.com/wilian746/go-generator/pkg/standart-gorm/configs"
	"github.com/wilian746/go-generator/pkg/standart-gorm/docs"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/routes"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/database"
	"log"
	"net/http"
)

// @title Standart Gorm
// @version 1.0
// @description This is a sample server using standart gorm server.
// @termsOfService http://swagger.io/terms/

// @contact.name Standart Gorm Support
// @contact.url https://github.com/wilian746/go-generator/issues
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/wilian746/go-generator/blob/master/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs := config.GetConfig()
	entity := &product.Product{}

	connection := database.GetConnection(configs.Dialect, configs.DatabaseURI)
	connection.Table(entity.TableName()).AutoMigrate(entity)
	repository := adapter.NewAdapter(connection)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	log.Println("service running on port ", port)

	setupSwagger()
	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func setupSwagger() {
	configs := config.GetConfig()
	// If your change host to Ex.: 192.168.1.0 is necessary change manually you field of search to your host too in your browser
	docs.SwaggerInfo.Host = configs.SwaggerHost
	docs.SwaggerInfo.BasePath = routes.BasePath
	log.Println("swagger running on url: ", fmt.Sprintf("http://%s/swagger/index.html", docs.SwaggerInfo.Host))
}
