package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/rest/model"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
)

func (store *Store) AddChurch(ctx context.Context, churchIn *model.ChurchIn) (church *models.ChurchProfile, errVal error) {
	return store.DAL.InsertChurch(ctx, churchIn)
}

func (store *Store) ActivateChurch(ctx context.Context, churchPubId *uuid.UUID, isActive bool) error {
	return store.DAL.UpdateChurchActivationStatus(ctx, churchPubId, isActive)
}
