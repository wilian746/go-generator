package files

type Files string
type NoGo string

const (
	CmdMain                                 Files = "cmd/main.go"
	ConfigsConfigs                          Files = "configs/configs.go"
	ConfigsConfigsTest                      Files = "configs/configs_test.go"
	InternalControllersProductProduct       Files = "internal/controllers/product/product.go"
	InternalControllersProductProductTest   Files = "internal/controllers/product/product_test.go"
	InternalEntitiesBase                    Files = "internal/entities/base.go"
	InternalEntitiesProductProduct          Files = "internal/entities/product/product.go"
	InternalHandlersInterface               Files = "internal/handlers/interface.go"
	InternalHandlersProductProduct          Files = "internal/handlers/product/product.go"
	InternalHandlersProductProductTest      Files = "internal/handlers/product/product_test.go"
	InternalHandlersHealthHealth            Files = "internal/handlers/health/health.go"
	InternalHandlersHealthHealthTest        Files = "internal/handlers/health/health_test.go"
	InternalRoutesConfig                    Files = "internal/routes/config.go"
	InternalRoutesConfigTest                Files = "internal/routes/config_test.go"
	InternalRoutesRoutes                    Files = "internal/routes/routes.go"
	InternalRoutesRoutesTest                Files = "internal/routes/routes_test.go"
	InternalRulesProductProduct             Files = "internal/rules/product/product_test.go"
	InternalRulesProductProductTest         Files = "internal/rules/product/product.go"
	InternalUtilsEnvironmentEnvironment     Files = "internal/utils/environment/environment.go"
	InternalUtilsEnvironmentEnvironmentTest Files = "internal/utils/environment/environment_test.go"
	InternalUtilsHTTPResponse               Files = "internal/utils/http/response.go"
	InternalUtilsLoggerLogger               Files = "internal/utils/logger/logger.go"
	InternalUtilsLoggerLoggerTest           Files = "internal/utils/logger/logger_test.go"
	PkgRepositoryAdapterAdapter             Files = "pkg/repository/adapter/adapter.go"
	PkgRepositoryAdapterAdapterTest         Files = "pkg/repository/adapter/adapter_test.go"
	PkgRepositoryAdapterAdapterMock         Files = "pkg/repository/adapter/adapter_mock.go"
	PkgRepositoryDatabaseDatabase           Files = "pkg/repository/database/database.go"
	PkgRepositoryDatabaseDatabaseTest       Files = "pkg/repository/database/database_test.go"
	PkgRepositoryEntitiesInterface          Files = "pkg/repository/entities/interface.go"
	PkgRepositoryResponseResponse           Files = "pkg/repository/response/response.go"
	PkgRepositoryResponseResponseTest       Files = "pkg/repository/response/response_test.go"
	GitIgnore                               NoGo  = ".gitignore"
	GolangCi                                NoGo  = ".golangci.yml"
	Makefile                                NoGo  = "Makefile"
	GoMod                                   NoGo  = "go.mod"
	GoSum                                   NoGo  = "go.sum"
)

func Values() []Files {
	return []Files{
		CmdMain,
		ConfigsConfigs,
		ConfigsConfigsTest,
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
		InternalUtilsEnvironmentEnvironment,
		InternalUtilsEnvironmentEnvironmentTest,
		InternalUtilsHTTPResponse,
		InternalUtilsLoggerLogger,
		InternalUtilsLoggerLoggerTest,
		PkgRepositoryAdapterAdapter,
		PkgRepositoryAdapterAdapterTest,
		PkgRepositoryAdapterAdapterMock,
		PkgRepositoryDatabaseDatabase,
		PkgRepositoryDatabaseDatabaseTest,
		PkgRepositoryEntitiesInterface,
		PkgRepositoryResponseResponse,
		PkgRepositoryResponseResponseTest,
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
