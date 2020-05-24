package files

type Files string
type NoGo string

const (
	CmdMain                               Files = "cmd/main.go"
	ConfigConfig                          Files = "config/config.go"
	ConfigConfigTest                      Files = "config/config_test.go"
	InternalControllersProductProduct     Files = "internal/controllers/product/product.go"
	InternalControllersProductProductTest Files = "internal/controllers/product/product_test.go"
	InternalEntitiesBase                  Files = "internal/entities/base.go"
	InternalEntitiesProductProduct        Files = "internal/entities/product/product.go"
	InternalHandlersInterface             Files = "internal/handlers/interface.go"
	InternalHandlersProductProduct        Files = "internal/handlers/product/product.go"
	InternalHandlersProductProductTest    Files = "internal/handlers/product/product_test.go"
	InternalHandlersHealthHealth          Files = "internal/handlers/health/health.go"
	InternalHandlersHealthHealthTest      Files = "internal/handlers/health/health_test.go"
	InternalRoutesConfig                  Files = "internal/routes/config.go"
	InternalRoutesConfigTest              Files = "internal/routes/config_test.go"
	InternalRoutesRoutes                  Files = "internal/routes/routes.go"
	InternalRoutesRoutesTest              Files = "internal/routes/routes_test.go"
	InternalRulesProductProduct           Files = "internal/rules/product/product_test.go"
	InternalRulesProductProductTest       Files = "internal/rules/product/product.go"
	RepositoryAdapterAdapter              Files = "repository/adapter/adapter.go"
	RepositoryAdapterAdapterTest          Files = "repository/adapter/adapter_test.go"
	RepositoryAdapterAdapterMock          Files = "repository/adapter/adapter_mock.go"
	RepositoryDatabaseDatabase            Files = "repository/database/database.go"
	RepositoryDatabaseDatabaseTest        Files = "repository/database/database_test.go"
	RepositoryEntitiesInterface           Files = "repository/entities/interface.go"
	RepositoryResponseResponse            Files = "repository/response/response.go"
	RepositoryResponseResponseTest        Files = "repository/response/response_test.go"
	UtilsEnvironmentEnvironment           Files = "utils/environment/environment.go"
	UtilsEnvironmentEnvironmentTest       Files = "utils/environment/environment_test.go"
	UtilsHTTPResponse                     Files = "utils/http/response.go"
	UtilsLoggerLogger                     Files = "utils/logger/logger.go"
	UtilsLoggerLoggerTest                 Files = "utils/logger/logger_test.go"
	GitIgnore                             NoGo  = ".gitignore"
	GolangCi                              NoGo  = ".golangci.yml"
	Makefile                              NoGo  = "Makefile"
	GoMod                                 NoGo  = "go.mod"
	GoSum                                 NoGo  = "go.sum"
)

func Values() []Files {
	return []Files{
		CmdMain,
		ConfigConfig,
		ConfigConfigTest,
		InternalControllersProductProduct,
		InternalControllersProductProductTest,
		InternalEntitiesBase,
		InternalEntitiesProductProduct,
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
		RepositoryAdapterAdapter,
		RepositoryAdapterAdapterTest,
		RepositoryAdapterAdapterMock,
		RepositoryDatabaseDatabase,
		RepositoryDatabaseDatabaseTest,
		RepositoryEntitiesInterface,
		RepositoryResponseResponse,
		RepositoryResponseResponseTest,
		UtilsEnvironmentEnvironment,
		UtilsEnvironmentEnvironmentTest,
		UtilsHTTPResponse,
		UtilsLoggerLogger,
		UtilsLoggerLoggerTest,
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
