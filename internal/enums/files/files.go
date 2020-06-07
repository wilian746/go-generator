package files

type Files string
type NoGo string

const (
	CmdMain                                    Files = "cmd/main.go"
	ConfigsConfigs                             Files = "configs/configs.go"
	ConfigsConfigsTest                         Files = "configs/configs_test.go"
	DeploymentsDockerCompose                   Files = "deployments/docker-compose.yaml"
	DocsDocs                                   Files = "docs/docs.go"
	DocsSwaggerJSON                            Files = "docs/swagger.json"
	DocsSwaggerYAML                            Files = "docs/swagger.yaml"
	InternalControllersProductProduct          Files = "internal/controllers/product/product.go"
	InternalControllersProductProductTest      Files = "internal/controllers/product/product_test.go"
	InternalEntitiesBase                       Files = "internal/entities/base.go"
	InternalEntitiesHealthSwaggerEntities      Files = "internal/entities/health/swagger_entities.go"
	InternalEntitiesProductProduct             Files = "internal/entities/product/product.go"
	InternalEntitiesProductSwaggerEntities     Files = "internal/entities/product/swagger_entities.go"
	InternalHandlersInterface                  Files = "internal/handlers/interface.go"
	InternalHandlersProductProduct             Files = "internal/handlers/product/product.go"
	InternalHandlersProductProductTest         Files = "internal/handlers/product/product_test.go"
	InternalHandlersHealthHealth               Files = "internal/handlers/health/health.go"
	InternalHandlersHealthHealthTest           Files = "internal/handlers/health/health_test.go"
	InternalRoutesConfig                       Files = "internal/routes/config.go"
	InternalRoutesConfigTest                   Files = "internal/routes/config_test.go"
	InternalRoutesRoutes                       Files = "internal/routes/routes.go"
	InternalRoutesRoutesTest                   Files = "internal/routes/routes_test.go"
	InternalRulesProductProduct                Files = "internal/rules/product/product_test.go"
	InternalRulesProductProductTest            Files = "internal/rules/product/product.go"
	InternalUtilsEnvironmentEnvironment        Files = "internal/utils/environment/environment.go"
	InternalUtilsEnvironmentEnvironmentTest    Files = "internal/utils/environment/environment_test.go"
	InternalUtilsHTTPResponse                  Files = "internal/utils/http/response.go"
	InternalUtilsLoggerLogger                  Files = "internal/utils/logger/logger.go"
	InternalUtilsLoggerLoggerTest              Files = "internal/utils/logger/logger_test.go"
	MigrationsMysqlCreateTableProductsUp       Files = "migrations/mysql/20200607175350_create_table_products.up.sql"
	MigrationsMysqlCreateTableProductsDown     Files = "migrations/mysql/20200607175350_create_table_products.down.sql"
	MigrationsPostgresCreateTableProductsUp    Files = "migrations/postgres/20200607175350_create_table_products.up.sql"
	MigrationsPostgresCreateTableProductsDown  Files = "migrations/postgres/20200607175350_create_table_products.down.sql"
	MigrationsSQLServerCreateTableProductsUp   Files = "migrations/sqlserver/20200607175350_create_table_products.up.sql"
	MigrationsSQLServerCreateTableProductsDown Files = "migrations/sqlserver/20200607175350_create_table_products.down.sql"
	PkgRepositoryAdapterAdapter                Files = "pkg/repository/adapter/adapter.go"
	PkgRepositoryAdapterAdapterTest            Files = "pkg/repository/adapter/adapter_test.go"
	PkgRepositoryAdapterAdapterMock            Files = "pkg/repository/adapter/adapter_mock.go"
	PkgRepositoryDatabaseDatabase              Files = "pkg/repository/database/database.go"
	PkgRepositoryDatabaseDatabaseTest          Files = "pkg/repository/database/database_test.go"
	PkgRepositoryEntitiesInterface             Files = "pkg/repository/entities/interface.go"
	PkgRepositoryEntitiesManyToMany            Files = "pkg/repository/entities/many_to_many.go"
	PkgRepositoryEntitiesOneToMany             Files = "pkg/repository/entities/one_to_many.go"
	PkgRepositoryEntitiesOneToOne              Files = "pkg/repository/entities/one_to_one.go"
	PkgRepositoryResponseResponse              Files = "pkg/repository/response/response.go"
	PkgRepositoryResponseResponseTest          Files = "pkg/repository/response/response_test.go"
	Readme                                     Files = "README.md"
)

const (
	GitIgnore NoGo = ".gitignore"
	GolangCi  NoGo = ".golangci.yml"
	Makefile  NoGo = "Makefile"
	GoMod     NoGo = "go.mod"
	GoSum     NoGo = "go.sum"
)

func Values() []Files {
	return []Files{
		CmdMain,
		ConfigsConfigs,
		ConfigsConfigsTest,
		DeploymentsDockerCompose,
		DocsDocs,
		DocsSwaggerJSON,
		DocsSwaggerYAML,
		InternalControllersProductProduct,
		InternalControllersProductProductTest,
		InternalEntitiesBase,
		InternalEntitiesProductProduct,
		InternalEntitiesHealthSwaggerEntities,
		InternalEntitiesProductSwaggerEntities,
		InternalHandlersInterface,
		InternalHandlersProductProduct,
		InternalHandlersProductProductTest,
		InternalHandlersHealthHealth,
		InternalHandlersHealthHealthTest,
		InternalRoutesConfig,
		InternalRoutesConfigTest,
		InternalRoutesRoutes,
		InternalRoutesRoutesTest,
		InternalRulesProductProduct,
		InternalRulesProductProductTest,
		InternalUtilsEnvironmentEnvironment,
		InternalUtilsEnvironmentEnvironmentTest,
		InternalUtilsHTTPResponse,
		InternalUtilsLoggerLogger,
		InternalUtilsLoggerLoggerTest,
		MigrationsMysqlCreateTableProductsUp,
		MigrationsMysqlCreateTableProductsDown,
		MigrationsPostgresCreateTableProductsUp,
		MigrationsPostgresCreateTableProductsDown,
		MigrationsSQLServerCreateTableProductsUp,
		MigrationsSQLServerCreateTableProductsDown,
		PkgRepositoryAdapterAdapter,
		PkgRepositoryAdapterAdapterTest,
		PkgRepositoryAdapterAdapterMock,
		PkgRepositoryDatabaseDatabase,
		PkgRepositoryDatabaseDatabaseTest,
		PkgRepositoryEntitiesInterface,
		PkgRepositoryEntitiesManyToMany,
		PkgRepositoryEntitiesOneToMany,
		PkgRepositoryEntitiesOneToOne,
		PkgRepositoryResponseResponse,
		PkgRepositoryResponseResponseTest,
		Readme,
	}
}

func ValuesNoGO() []NoGo {
	return []NoGo{
		GitIgnore,
		GolangCi,
		Makefile,
		GoMod,
		GoSum,
	}
}
