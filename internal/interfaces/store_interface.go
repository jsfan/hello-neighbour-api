package interfaces

//go:generate go run github.com/vektra/mockery/v2 --name DataInterface --structname Store --case underscore

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

type DataInterface interface {
	Clone() DataInterface

	ActivateChurch(ctx context.Context, churchPubId *uuid.UUID, isActive bool) error
	AddChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error)
	DeleteUser(ctx context.Context, userPubId *uuid.UUID) error
	GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error)
	Migrate(dbName *string) (errVal error)
	PromoteToLeader(ctx context.Context, userPubId *uuid.UUID, churchPubId *uuid.UUID) error
	RegisterUser(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error)
}
