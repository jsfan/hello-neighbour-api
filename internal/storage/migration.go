package storage

// Migrates the database schema to the latest schema version
func (store *Store) Migrate(dbName *string) (errVal error) {
	dalInstance, commitFunc, rollbackFunc, err := store.GetDAL(nil)
	if err != nil {
		return err
	}
	err = dalInstance.Migrate(dbName)
	if err != nil {
		return err
	}
	err = commitFunc()
	if err != nil {
		rollbackFunc()
		return err
	}
	return nil
}
