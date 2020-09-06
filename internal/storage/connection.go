package storage

import (
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/pkg/errors"
)

var backend *Store

func Connect(dbConfig *config.DatabaseConfig) (connection *Store, errVal error) {
	dalInstance, err := dal.Connect(dbConfig)
	if err != nil {
		return nil, err
	}
	backend = &Store{
		dal: dalInstance,
	}
	return backend, nil
}

func GetStore() (conn *Store, errVal error) {
	if backend == nil {
		return nil, errors.New("No database connection.")
	}
	return backend, nil
}
