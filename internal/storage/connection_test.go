package storage_test

import (
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/storage"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces/mocks"
)

func ConnectMock(dbConfig *config.DatabaseConfig) (connection *storage.Store) {
	return &storage.Store{
		DAL: &mocks.DAL{},
	}
}
