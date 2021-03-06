package dal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jsfan/hello-neighbour-api/internal/config"
)

func Connect(dbConfig *config.DatabaseConfig) (connection AccessInterface, errVal error) {
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

func (dalInstance *DAL) SetupDal(ctx context.Context) (commit func() error, rollback func() error, errVal error) {
	dalInstance.ctx = ctx
	if dalInstance.tx == nil && dalInstance.ctx != nil {
		var err error
		dalInstance.tx, err = dalInstance.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return func() error {
			defer func() {
				dalInstance.tx = nil
			}()
			return dalInstance.tx.Commit()
		}, func() error {
			defer func() {
				dalInstance.tx = nil
			}()
			return dalInstance.tx.Rollback()
		}, nil
}
