package storage

import (
	"github.com/jsfan/hello-neighbour/internal/config"
)

type DBInteraction interface {
	Connect(dbConfig *config.DatabaseConfig) (conn *DBConnection, errVal error)
	GetUserByEmail(username string)
}
