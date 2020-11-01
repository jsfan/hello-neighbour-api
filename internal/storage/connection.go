package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/pkg/errors"
)

var backend *Store

func Connect(dbConfig *config.DatabaseConfig) (connection DataInterface, errVal error) {
	dalInstance, err := dal.Connect(dbConfig)
	if err != nil {
		return nil, err
	}
	backend = &Store{
		DAL: dalInstance,
	}
	return backend, nil
}

func GetStore() (conn *Store, errVal error) {
	if backend == nil {
		return nil, errors.New("No database connection.")
	}
	return backend, nil
}

func (store *Store) GetDAL(ctx context.Context) (dalInstance dal.AccessInterface, commitFunc func() error, errVal error) {
	commitFunc, err := store.DAL.SetupDal(ctx)
	if err != nil {
		return nil, nil, err
	}
	return store.DAL, commitFunc, nil
}

func setupContext(ctx context.Context) (ctext context.Context, cancelCtx context.CancelFunc) {
	ctext, cancelCtx = context.WithCancel(ctx)
	return ctext, cancelCtx
}
