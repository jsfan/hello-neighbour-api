package storage

// Migrate migrates the database schema to the latest schema version
func (store *Store) Migrate(dbName *string) (errVal error) {
	return store.DAL.Migrate(dbName)
}
