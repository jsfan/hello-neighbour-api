package storage_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/storage/dal"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
	"reflect"
	"testing"
)

func TestStore_GetUserByEmail(t *testing.T) {
	// TODO: Add error cases.
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
		"SetupDAL":          dal.ResponseSignature{{func() error { return nil }, func() error { return nil }, nil}},
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

	// there should be two calls
	if len(mDAL.Calls) != 2 {
		t.Errorf("Unexpected number of calls to DAL. Expected %d, got %d.", 2, len(mDAL.Calls))
	}

	// first call should be to SetupDal
	expectedFunction := "SetupDal"
	if mDAL.Calls[0].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %s, got %s.", expectedFunction, mDAL.Calls[0].FunctionName)
	}

	// second call should be to GetUserByEmail
	expectedFunction = "SelectUserByEmail"
	expectedParams := []string{"test@example.com"}
	if mDAL.Calls[1].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[1].Args, expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[1].Args)
	}
}

func TestStore_RegisterUser(t *testing.T) {
	// TODO: Add error cases.
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
		"SetupDAL":          dal.ResponseSignature{{func() error { return nil }, func() error { return nil }, nil}},
		"InsertUser":        dal.ResponseSignature{{nil}},
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

	// there should be three calls
	if len(mDAL.Calls) != 3 {
		t.Errorf("Unexpected number of calls to DAL. Expected %d, got %d.", 3, len(mDAL.Calls))
	}

	// first call should be to SetupDal
	expectedFunction := "SetupDal"
	if mDAL.Calls[0].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %s, got %s.", expectedFunction, mDAL.Calls[0].FunctionName)
	}

	// second call should be to InsertUser
	expectedFunction = "InsertUser"
	if mDAL.Calls[1].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[1].Args, userIn) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", userIn, mDAL.Calls[1].Args)
	}

	// third call should be to SelectUserByEmail
	expectedFunction = "SelectUserByEmail"
	expectedParams := []string{"test@example.com"}
	if mDAL.Calls[2].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[2].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[2].Args, expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[2].Args)
	}
}

func TestStore_DeleteUser(t *testing.T) {
	// TODO: Add error cases.
	store := ConnectMock(&config.DatabaseConfig{})
	ctx := context.Background()

	mDAL := store.DAL.(*dal.MockDAL)
	mDAL.Responses = dal.ResponseMap{
		"SetupDAL":          dal.ResponseSignature{{func() error { return nil }, func() error { return nil }, nil}},
		"DeleteUserByPubId": dal.ResponseSignature{{nil}},
	}

	expectedUUID := uuid.New()
	err := store.DeleteUser(ctx, &expectedUUID)

	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	// there should be two calls
	if len(mDAL.Calls) != 2 {
		t.Errorf("Unexpected number of calls to DAL. Expected %d, got %d.", 2, len(mDAL.Calls))
	}

	// first call should be to SetupDal
	expectedFunction := "SetupDal"
	if mDAL.Calls[0].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %s, got %s.", expectedFunction, mDAL.Calls[0].FunctionName)
	}

	// second call should be to DeleteUserByPubId
	expectedFunction = "DeleteUserByPubId"
	expectedParams := []interface{}{&uuid.UUID{}}
	if mDAL.Calls[1].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[1].Args, expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[1].Args)
	}
}
