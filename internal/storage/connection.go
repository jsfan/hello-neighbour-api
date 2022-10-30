package storage

import (
	"github.com/jsfan/hello-neighbour-api/internal/config"
	storeInterface "github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"github.com/jsfan/hello-neighbour-api/internal/storage/dal"
	dalInterface "github.com/jsfan/hello-neighbour-api/internal/storage/interfaces"
)

type Store struct {
	DAL dalInterface.AccessInterface
}

func Connect(dbConfig *config.DatabaseConfig) (connection storeInterface.DataInterface, errVal error) {
	dalInstance, err := dal.Connect(dbConfig)
	if err != nil {
		return nil, err
	}
	return &Store{
		DAL: dalInstance,
	}, nil
}

func (store *Store) Clone() storeInterface.DataInterface {
	newDAL := store.DAL.Clone()
	if newDAL == nil {
		return nil
	}
	return &Store{DAL: newDAL}
}
