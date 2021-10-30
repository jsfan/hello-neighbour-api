package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

type CallSignature struct {
	FunctionName string
	Args         []interface{}
}

type ResponseSignature [][]interface{}

type ResponseMap map[string]ResponseSignature

type MockDAL struct {
	Calls     []*CallSignature
	Responses ResponseMap
	txOn bool
}

func addCall(mDAL *MockDAL, functionName string, args ...interface{}) {
	if mDAL.Calls == nil {
		mDAL.Calls = []*CallSignature{{
			FunctionName: functionName,
			Args:         args,
		}}
	} else {
		mDAL.Calls = append(
			mDAL.Calls,
			&CallSignature{
				FunctionName: functionName,
				Args:         args,
			},
		)
	}
}

func getResponse(mDAL *MockDAL, functionName string) []interface{} {
	response := mDAL.Responses[functionName][0]
	mDAL.Responses[functionName] = mDAL.Responses[functionName][1:]
	return response
}

func castError(rawError interface{}) error {
	typedError, _ := rawError.(error)
	return typedError
}

func (mDAL *MockDAL) Clone() AccessInterface {
	addCall(mDAL, "Clone")
	return mDAL
}

func (mDAL *MockDAL) BeginTransaction() error {
	addCall(mDAL, "BeginTransaction")
	response := getResponse(mDAL, "BeginTransaction")
	return castError(response[0])
}

func (mDAL *MockDAL) CancelTransaction() error {
	addCall(mDAL, "CancelTransaction")
	response := getResponse(mDAL, "CancelTransaction")
	return castError(response[0])
}

func (mDAL *MockDAL) CompleteTransaction() error {
	addCall(mDAL, "CompleteTransaction")
	response := getResponse(mDAL, "CompleteTransaction")
	return castError(response[0])
}

func (mDAL *MockDAL) DeleteUserByPubId(ctx context.Context, userPubId *uuid.UUID) error {
	addCall(mDAL, "DeleteUserByPubId", ctx, userPubId)
	response := getResponse(mDAL, "DeleteUserByPubId")
	return castError(response[0])
}

func (mDAL *MockDAL) InsertChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error) {
	addCall(mDAL, "InsertChurch", ctx, churchIn)
	response := getResponse(mDAL, "InsertChurch")
	return response[0].(*models.ChurchProfile), castError(response[1])
}

func (mDAL *MockDAL) InsertUser(ctx context.Context, userIn *pkg.UserIn) error {
	addCall(mDAL, "InsertUser", ctx, userIn)
	response := getResponse(mDAL, "InsertUser")
	return castError(response[0])
}

func (mDAL *MockDAL) MakeLeader(ctx context.Context, userPubId *uuid.UUID, churchPubId *uuid.UUID) error {
	addCall(mDAL, "MakeLeader", ctx, userPubId, churchPubId)
	response := getResponse(mDAL, "MakeLeader")
	return castError(response[0])
}

func (mDAL *MockDAL) Migrate(dbName *string) (errVal error) {
	addCall(mDAL, "Migrate", dbName)
	response := getResponse(mDAL, "Migrate")
	return castError(response[0])
}

func (mDAL *MockDAL) SelectChurchByEmail(ctx context.Context, email string) (church *models.ChurchProfile, errVal error) {
	addCall(mDAL, "SelectChurchByEmail", ctx, email)
	response := getResponse(mDAL, "SelectChurchByEmail")
	return response[0].(*models.ChurchProfile), castError(response[1])
}

func (mDAL *MockDAL) SelectUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	addCall(mDAL, "SelectUserByEmail", ctx, email)
	response := getResponse(mDAL, "SelectUserByEmail")
	return response[0].(*models.UserProfile), castError(response[1])
}

func (mDAL *MockDAL) UpdateChurchActivationStatus(ctx context.Context, churchPubId *uuid.UUID, isActive bool) error {
	addCall(mDAL, "UpdateChurchActivationStatus", ctx, churchPubId, isActive)
	response := getResponse(mDAL, "UpdateChurchActivationStatus")
	return castError(response[0])
}
