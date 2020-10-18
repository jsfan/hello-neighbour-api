package storage_test

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
	"testing"
)

func TestStore_GetUserByEmail(t *testing.T) {
	store := ConnectMock(&config.DatabaseConfig{})
	ctx := context.Background()

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

	mDAL := store.DAL.(*dal.MockDAL)
	mDAL.Responses = dal.ResponseMap{
		"SetupDAL":          dal.ResponseSignature{{func() error { return nil }, nil}},
		"SelectUserByEmail": dal.ResponseSignature{{expectedUser, nil}},
	}

	user, err := store.GetUserByEmail(ctx, "test@example.com")

	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	if user != expectedUser {
		t.Errorf("User record differs from expected record. Expected %+v, got %+v", expectedUser, user)
	}
}

func TestStore_UserRegister(t *testing.T) {
	store := ConnectMock(&config.DatabaseConfig{})
	ctx := context.Background()

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

	mDAL := store.DAL.(*dal.MockDAL)
	mDAL.Responses = dal.ResponseMap{
		"SetupDAL":          dal.ResponseSignature{{func() error { return nil }, nil}},
		"InsertUser":      dal.ResponseSignature{{nil}},
		"SelectUserByEmail": dal.ResponseSignature{{expectedUser, nil}},
	}

	user, err := store.RegisterUser(ctx, userIn)

	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	if user != expectedUser {
		t.Errorf("User record differs from expected record. Expected %+v, got %+v", expectedUser, user)
	}
}
