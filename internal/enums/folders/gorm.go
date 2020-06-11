package folders

const (
	MigrationsMysql     Folders = "migrations/mysql"
	MigrationsPostgres  Folders = "migrations/postgres"
	MigrationsSQLServer Folders = "migrations/sqlserver"
)

func ValuesGorm() []Folders {
	return []Folders{
		MigrationsMysql,
		MigrationsPostgres,
		MigrationsSQLServer,
	}
}
