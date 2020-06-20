package storage

import (
	"database/sql"
	"fmt"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/pkg/errors"
)

type DBConnection struct {
	Db *sql.DB
}

var backend *DBConnection

func (conn *DBConnection) Connect(dbConfig *config.DatabaseConfig) (connection *DBConnection, errVal error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	database, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	conn.Db = database
	return backend, nil
}

func GetConnection() (conn *DBConnection, errVal error) {
	if backend == nil {
		return nil, errors.New("No database connection.")
	}
	return backend, nil
}
