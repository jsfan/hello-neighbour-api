package dal

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // blank import to register source driver
)

// Migrate migrates the database schema to the latest schema version
func (dalInstance *DAL) Migrate(dbName *string) (errVal error) {
	driver, err := postgres.WithInstance(dalInstance.Db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", *dbName, driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
