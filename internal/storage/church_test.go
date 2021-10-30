package storage_test

import (
	"context"
	"github.com/jsfan/hello-neighbour-api/internal/storage/interfaces"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
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

	mDAL := store.DAL.(*interfaces.MockDAL)
	mDAL.Responses = interfaces.ResponseMap{
		"InsertChurch": interfaces.ResponseSignature{{expectedChurch, nil}},
	}

	church, err := store.AddChurch(ctx, churchIn)

	// no error expected
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	if church != expectedChurch {
		t.Errorf("Church record differs from expected record. Expected %+v, got %+v", expectedChurch, church)
	}

	// there should be one call
	const expectedCalls = 1
	if len(mDAL.Calls) != expectedCalls {
		t.Errorf("Unexpected number of calls to DAL. Expected %d, got %d.", expectedCalls, len(mDAL.Calls))
	}

	// call should be to InsertChurch
	expectedFunction := "InsertChurch"
	if mDAL.Calls[0].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[0].Args, churchIn) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", churchIn, mDAL.Calls[1].Args)
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

	mDAL := store.DAL.(*interfaces.MockDAL)
	mDAL.Responses = interfaces.ResponseMap{
		"UpdateChurchActivationStatus": interfaces.ResponseSignature{{&churchPubId, isActive}},
	}

	err = store.ActivateChurch(ctx, &churchPubId, isActive)
	if err != nil {
		t.Errorf("Got unexpected error: %+v", err)
	}

	// there should be one call
	const expectedCalls = 1
	if len(mDAL.Calls) != expectedCalls {
		t.Errorf("Unexpected number of calls to DAL. Expected %d, got %d.", expectedCalls, len(mDAL.Calls))
	}

	// call should be to UpdateChurchActivationStatus
	expectedFunction := "UpdateChurchActivationStatus"
	expectedParams := []interface{}{&churchPubId, isActive}
	if mDAL.Calls[0].FunctionName != expectedFunction {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedFunction, mDAL.Calls[1].FunctionName)
	}
	if reflect.DeepEqual(mDAL.Calls[0], expectedParams) {
		t.Errorf("Recorded call not as expected. Expected function %+v, got %+v.", expectedParams, mDAL.Calls[1].Args)
	}
}
