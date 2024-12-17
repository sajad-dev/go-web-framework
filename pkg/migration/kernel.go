package migration

var ArrMigrations []*Migrate

func MigrationList(migrate []*Migrate) {
	ArrMigrations = append(ArrMigrations, migrate...)
}
