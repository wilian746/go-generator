package folders

type Folders string

const (
	Cmd                        Folders = "cmd"
	Config                     Folders = "config"
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
	Repository                 Folders = "repository"
	RepositoryAdapter          Folders = "repository/adapter"
	RepositoryDatabase         Folders = "repository/database"
	RepositoryEntities         Folders = "repository/entities"
	RepositoryResponse         Folders = "repository/response"
	Utils                      Folders = "utils"
	UtilsEnvironment           Folders = "utils/environment"
	UtilsHTTP                  Folders = "utils/http"
	UtilsLogger                Folders = "utils/logger"
)

func Values() []Folders {
	return []Folders{
		Cmd,
		Config,
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
		Repository,
		RepositoryAdapter,
		RepositoryDatabase,
		RepositoryEntities,
		RepositoryResponse,
		Utils,
		UtilsEnvironment,
		UtilsHTTP,
		UtilsLogger,
	}
}
