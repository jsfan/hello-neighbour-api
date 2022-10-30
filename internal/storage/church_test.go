package storage_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/rest/model"
	"github.com/jsfan/hello-neighbour-api/internal/storage"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces/mocks"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStore_AddChurch(t *testing.T) {
	// TODO: Add error cases
	churchIn := &model.ChurchIn{
		Name:                  "Test Church",
		Description:           "description",
		Address:               "Church Avenue",
		Website:               "church.com",
		Email:                 "church_email@church.com",
		Phone:                 "777-7777",
		GroupSize:             2,
		SameGender:            true,
		MinAge:                13,
		MemberBasicInfoUpdate: false,
	}

	expectedChurch := &models.ChurchProfile{
		PubId:                 "3da2eb8d-d9d1-44bf-bdd4-7fb0a83f2f77",
		Name:                  "Test Church",
		Description:           "description",
		Address:               "Church Avenue",
		Website:               "church.com",
		Email:                 "church_email@church.com",
		Phone:                 "777-7777",
		GroupSize:             "2",
		SameGender:            true,
		MinAge:                "13",
		MemberBasicInfoUpdate: false,
		Active:                false,
	}

	ctx := context.Background()

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("InsertChurch", ctx, churchIn).
		Return(expectedChurch, nil).
		Run(func(args mock.Arguments) {
			ctxRcv := args.Get(0).(context.Context)
			churchInRv := args.Get(1).(*model.ChurchIn)
			assert.Equal(t, ctxRcv, ctx)
			assert.Equal(t, churchInRv, churchIn)
		})

	store := &storage.Store{
		DAL: dalMock,
	}

	church, err := store.AddChurch(ctx, churchIn)

	require.Equal(t, expectedChurch, church)
	require.Nil(t, err)

	dalMock.AssertExpectations(t)
}

func TestStore_ActivateChurch(t *testing.T) {
	// TODO: add error cases
	ctx := context.Background()
	churchPubId := uuid.New()
	const isActiveFlag = true

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("UpdateChurchActivationStatus", ctx, &churchPubId, isActiveFlag).
		Return(nil).
		Run(func(args mock.Arguments) {
			ctxRcv := args.Get(0).(context.Context)
			churchPubIdRcv := args.Get(1).(*uuid.UUID)
			isActiveRcv := args.Get(2).(bool)
			assert.Equal(t, ctxRcv, ctx)
			assert.Equal(t, *churchPubIdRcv, churchPubId)
			assert.Equal(t, isActiveRcv, isActiveFlag)
		})

	store := &storage.Store{
		DAL: dalMock,
	}

	err := store.ActivateChurch(ctx, &churchPubId, true)

	require.Nil(t, err)
	dalMock.AssertExpectations(t)
}
