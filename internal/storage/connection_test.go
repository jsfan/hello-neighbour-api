package storage_test

import (
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
)

func ConnectMock(dbConfig *config.DatabaseConfig) (connection *storage.Store) {
	dalInstance := &dal.MockDAL{}
	return &storage.Store{
		DAL: dalInstance,
	}
}