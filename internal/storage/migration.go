package storage

import (
	"github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file" // blank import to register source driver
)

// Migrates the database schema to the latest schema version
func Migrate(connection *DBConnection) (errVal error) {
	driver, err := postgres.WithInstance(connection.Db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
    if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
