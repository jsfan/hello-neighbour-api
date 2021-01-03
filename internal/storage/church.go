package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

func (store *Store) AddChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error) {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, rollbackFunc, err := store.GetDAL(ctx)
	if err != nil {
		rollbackFunc()
		cancelCtx()
		return nil, err
	}
	if err = dbAccess.InsertChurch(churchIn); err != nil {
		rollbackFunc()
		cancelCtx()
		return nil, err
	}
	church, err = dbAccess.SelectChurchByEmail(churchIn.Email)
	if err != nil {
		return nil, err
	}
	if err = commitFunc(); err != nil {
		rollbackFunc()
		cancelCtx()
		return nil, err
	}
	return church, nil
}

func (store *Store) ChurchActivation(ctx context.Context, churchPubId *uuid.UUID, isActive bool) error {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, rollbackFunc, err := store.GetDAL(ctx)
	if err != nil {
		rollbackFunc()
		cancelCtx()
		return err
	}
	if err = dbAccess.UpdateChurchActivationStatus(churchPubId, isActive); err != nil {
		rollbackFunc()
		cancelCtx()
		return err
	}
	if err = commitFunc(); err != nil {
		rollbackFunc()
		cancelCtx()
		return err
	}
	return nil
}
