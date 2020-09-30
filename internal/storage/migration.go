package storage

import (
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
)

// Migrates the database schema to the latest schema version
func (store *Store) Migrate(dbName *string) (errVal error) {
	dalInstance, commitFunc, err := dal.GetDAL(nil)
	if err != nil {
		return err
	}
	err = dalInstance.Migrate(dbName)
	if err != nil {
		return err
	}
	err = commitFunc()
	if err != nil {
		return err
	}
	return nil
}
