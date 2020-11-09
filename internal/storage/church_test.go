package storage_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
	"reflect"
	"testing"
)

func TestStore_AddChurch(t *testing.T) {
	// TODO: Add error cases
	store := ConnectMock(&config.DatabaseConfig{})
	ctx := context.Background()

	churchIn := &pkg.ChurchIn{
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

	mDAL := store.DAL.(*dal.MockDAL)
	mDAL.Responses = dal.ResponseMap{
		"SetupDAL":            dal.ResponseSignature{{func() error { return nil }, nil}},
		"InsertChurch":        dal.ResponseSignature{{nil}},
		"SelectChurchByEmail": dal.ResponseSignature{{expectedChurch, nil}},
	}

	church, err := store.AddChurch(ctx, churchIn)

	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	if church != expectedChurch {
		t.Errorf("Church record differs from expected record. Expected %+v, got %+v", expectedChurch, church)
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

	// second call should be to InsertChurch
	expectedFunction = "InsertChurch"
	if mDAL.Calls[1].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[1].Args, churchIn) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", churchIn, mDAL.Calls[1].Args)
	}

	// third call should be to SelectChurchByEmail
	expectedFunction = "SelectChurchByEmail"
	expectedParams := []string{"church_email@church.com"}
	if mDAL.Calls[2].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[2].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[2].Args, expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[2].Args)
	}
}

func TestStore_ChurchActivation(t *testing.T) {
	// TODO: add error cases
	store := ConnectMock(&config.DatabaseConfig{})
	ctx := context.Background()

	churchPubId, err := uuid.Parse("3da2eb8d-d9d1-44bf-bdd4-7fb0a83f2f77")
	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}
	isActive := true

	mDAL := store.DAL.(*dal.MockDAL)
	mDAL.Responses = dal.ResponseMap{
		"SetupDAL":            dal.ResponseSignature{{func() error { return nil }, nil}},
		"UpdateChurchActivationStatus":        dal.ResponseSignature{{&churchPubId, isActive}},
	}

	err = store.ChurchActivation(ctx, &churchPubId, isActive)
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

	// second call should be to UpdateChurchActivationStatus
	expectedFunction = "UpdateChurchActivationStatus"
	expectedParams := []interface{}{&churchPubId, isActive}
	if mDAL.Calls[1].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[1], expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[1].Args)
	}
}
