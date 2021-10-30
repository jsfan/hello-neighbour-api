package storage_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces/mocks"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_GetUserByEmail(t *testing.T) {
	// TODO: Add error cases.

	expectedDob := "1964-01-07"
	expectedDescription := "A Description"
	expectedUser := &models.UserProfile{
		PubId:        "8fcf0e39-4885-4f4d-bc87-64de887c40fc",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$0dG1FfR6YZ5fbHOoIukIY.3cm2oR0F3zoBhzAhN1YYMuZ06vqJg4e",
		FirstName:    "Test",
		LastName:     "User",
		DateOfBirth:  &expectedDob,
		Gender:       "male",
		Description:  &expectedDescription,
		ChurchUUID:   "0dfc7330-2b8d-4391-96a9-059332ef9874",
		Role:         "member",
		Active:       true,
	}

	ctx := context.Background()

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("SelectUserByEmail", ctx, expectedUser.Email).
		Return(expectedUser, nil).
		Run(func(args mock.Arguments) {
			ctxRcv := args.Get(0).(context.Context)
			emailRcv := args.Get(1).(string)
			assert.Equal(t, ctxRcv, ctx)
			assert.Equal(t, emailRcv, expectedUser.Email)
		})

	store := &storage.Store{
		DAL: dalMock,
	}

	user, err := store.GetUserByEmail(ctx, expectedUser.Email)

	require.Equal(t, expectedUser, user)
	require.Nil(t, err)

	dalMock.AssertExpectations(t)
}

func TestStore_RegisterUser(t *testing.T) {
	// TODO: Add error cases.

	expectedDob := "1964-01-07"
	expectedDescription := "A Description"
	userIn := &pkg.UserIn{
		Email:       "test@example.com",
		FirstName:   "Test",
		LastName:    "User",
		Gender:      "male",
		Description: expectedDescription,
		Church:      "0dfc7330-2b8d-4391-96a9-059332ef9874",
		Role:        "member",
		DateOfBirth: expectedDob,
		Password:    "NotHashedHere",
	}

	expectedUser := &models.UserProfile{
		PubId:        "8fcf0e39-4885-4f4d-bc87-64de887c40fc",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$0dG1FfR6YZ5fbHOoIukIY.3cm2oR0F3zoBhzAhN1YYMuZ06vqJg4e",
		FirstName:    "Test",
		LastName:     "User",
		DateOfBirth:  &expectedDob,
		Gender:       "male",
		Description:  &expectedDescription,
		ChurchUUID:   "0dfc7330-2b8d-4391-96a9-059332ef9874",
		Role:         "member",
		Active:       true,
	}

	ctx := context.Background()

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("BeginTransaction").
		Return(nil)
	dalMock.
		On("InsertUser", ctx, userIn).
		Return(nil).
		Run(func(args mock.Arguments) {
			ctxRcv := args.Get(0).(context.Context)
			userRcv := args.Get(1).(*pkg.UserIn)
			assert.Equal(t, ctxRcv, ctx)
			assert.Equal(t, userRcv, userIn)
		})
	dalMock.
		On("SelectUserByEmail", ctx, expectedUser.Email).
		Return(expectedUser, nil).
		Run(func(args mock.Arguments) {
			ctxRcv := args.Get(0).(context.Context)
			emailRcv := args.Get(1).(string)
			assert.Equal(t, ctxRcv, ctx)
			assert.Equal(t, emailRcv, expectedUser.Email)
		})
	dalMock.
		On("CompleteTransaction").
		Return(nil)

	store := &storage.Store{
		DAL: dalMock,
	}

	user, err := store.RegisterUser(ctx, userIn)

	require.Equal(t, expectedUser, user)
	require.Nil(t, err)

	dalMock.AssertExpectations(t)
}

func TestStore_DeleteUser(t *testing.T) {
	// TODO: Add error cases.
	ctx := context.Background()
	userUUID := uuid.New()

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("DeleteUserByPubId", ctx, &userUUID).
		Return(nil)

	store := &storage.Store{
		DAL: dalMock,
	}

	err := store.DeleteUser(ctx, &userUUID)

	require.Nil(t, err)
}

func TestStore_PromoteToLeader(t *testing.T) {
	// TODO: add error cases
	ctx := context.Background()
	userUUID := uuid.New()
	churchUUID := uuid.New()

	// prepare DAL mock
	dalMock := &mocks.DAL{}
	dalMock.
		On("MakeLeader", ctx, &churchUUID, &userUUID).
		Return(nil)

	store := &storage.Store{
		DAL: dalMock,
	}

	err := store.PromoteToLeader(ctx, &userUUID, &churchUUID)

	require.Nil(t, err)
}
