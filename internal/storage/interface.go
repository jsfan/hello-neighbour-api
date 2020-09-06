package storage

import "github.com/jsfan/hello-neighbour/internal/storage/dal"

type Store struct {
	dal dal.AccessInterface
}

type DataInterface interface {
	Migrate(dbName *string) (errVal error)
	GetUserByEmail(username string)
}
