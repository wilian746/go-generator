package folders

type Folders string

const (
	Cmd                        Folders = "cmd"
	Configs                    Folders = "configs"
	Internal                   Folders = "internal"
	InternalControllers        Folders = "internal/controllers"
	InternalControllersProduct Folders = "internal/controllers/product"
	InternalEntities           Folders = "internal/entities"
	InternalEntitiesProduct    Folders = "internal/entities/product"
	InternalHandlers           Folders = "internal/handlers"
	InternalHandlersProduct    Folders = "internal/handlers/product"
	InternalHandlersHealth     Folders = "internal/handlers/health"
	InternalRoutes             Folders = "internal/routes"
	InternalRules              Folders = "internal/rules"
	InternalRulesProduct       Folders = "internal/rules/product"
	InternalUtils              Folders = "internal/utils"
	InternalUtilsEnvironment   Folders = "internal/utils/environment"
	InternalUtilsHTTP          Folders = "internal/utils/http"
	InternalUtilsLogger        Folders = "internal/utils/logger"
	Pkg                        Folders = "pkg/repository"
	PkgRepository              Folders = "pkg/repository"
	PkgRepositoryAdapter       Folders = "pkg/repository/adapter"
	PkgRepositoryDatabase      Folders = "pkg/repository/database"
	PkgRepositoryEntities      Folders = "pkg/repository/entities"
	PkgRepositoryResponse      Folders = "pkg/repository/response"
)

func Values() []Folders {
	return []Folders{
		Cmd,
		Configs,
		Internal,
		InternalControllers,
		InternalControllersProduct,
		InternalEntities,
		InternalEntitiesProduct,
		InternalHandlers,
		InternalHandlersProduct,
		InternalHandlersHealth,
		InternalRoutes,
		InternalRules,
		InternalRulesProduct,
		Pkg,
		PkgRepository,
		PkgRepositoryAdapter,
		PkgRepositoryDatabase,
		PkgRepositoryEntities,
		PkgRepositoryResponse,
		InternalUtils,
		InternalUtilsEnvironment,
		InternalUtilsHTTP,
		InternalUtilsLogger,
	}
}
