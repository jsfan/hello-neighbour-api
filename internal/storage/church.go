package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

// AddChurch adds a new church
func (store *Store) AddChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error) {
	// TODO: Implement me
	return &models.ChurchProfile{}, nil
}
