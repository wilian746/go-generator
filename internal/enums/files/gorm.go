package files

const (
	MigrationsMysqlCreateTableProductsUp       Files = "migrations/mysql/20200607175350_create_table_products.up.sql"
	MigrationsMysqlCreateTableProductsDown     Files = "migrations/mysql/20200607175350_create_table_products.down.sql"
	MigrationsPostgresCreateTableProductsUp    Files = "migrations/postgres/20200607175350_create_table_products.up.sql"
	MigrationsPostgresCreateTableProductsDown  Files = "migrations/postgres/20200607175350_create_table_products.down.sql"
	MigrationsSQLServerCreateTableProductsUp   Files = "migrations/sqlserver/20200607175350_create_table_products.up.sql"
	MigrationsSQLServerCreateTableProductsDown Files = "migrations/sqlserver/20200607175350_create_table_products.down.sql"
)

func ValuesGorm() []Files {
	return []Files{
		MigrationsMysqlCreateTableProductsUp,
		MigrationsMysqlCreateTableProductsDown,
		MigrationsPostgresCreateTableProductsUp,
		MigrationsPostgresCreateTableProductsDown,
		MigrationsSQLServerCreateTableProductsUp,
		MigrationsSQLServerCreateTableProductsDown,
	}
}
