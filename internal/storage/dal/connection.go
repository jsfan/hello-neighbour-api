package dal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/pkg/errors"
)

var dal *DAL

func Connect(dbConfig *config.DatabaseConfig) (connection *DAL, errVal error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	database, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	dal = &DAL{
		Db: database,
	}
	return dal, nil
}

func GetDAL(ctx context.Context) (conn *DAL, commit func() error, errVal error) {
	if dal == nil {
		return nil, nil, errors.New("No database connection.")
	}
	dal.ctx = ctx
	if dal.tx == nil && dal.ctx != nil {
		var err error
		dal.tx, err = dal.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return dal, func() error { return dal.tx.Commit() },nil
}
