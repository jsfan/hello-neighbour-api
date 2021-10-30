package storage_test

import (
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/storage"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces"
)

func ConnectMock(dbConfig *config.DatabaseConfig) (connection *storage.Store) {
	dalInstance := &interfaces.MockDAL{}
	return &storage.Store{
		DAL: dalInstance,
	}
}
