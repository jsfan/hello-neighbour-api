package dal

import (
	"database/sql"
	"fmt"

	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces"
)

type DAL struct {
	Db *sql.DB
	tx *sql.Tx
}

func Connect(dbConfig *config.DatabaseConfig) (connection interfaces.AccessInterface, errVal error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	database, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	dalInstance := &DAL{
		Db: database,
	}
	return dalInstance, nil
}

func (dalInstance *DAL) Clone() interfaces.AccessInterface {
	return &DAL{
		Db: dalInstance.Db,
		tx: nil,
	}
}

func (dalInstance *DAL) db() interfaces.DbInteractionInterface {
	if dalInstance.tx != nil {
		return dalInstance.tx
	}
	return dalInstance.Db
}
