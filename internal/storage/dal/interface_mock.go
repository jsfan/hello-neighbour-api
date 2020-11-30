package dal

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
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

func (mDAL *MockDAL) SetupDal(ctx context.Context) (commit func() error, rollback func() error, errVal error) {
	addCall(mDAL, "SetupDal", ctx)
	response := getResponse(mDAL, "SetupDAL")
	return response[0].(func() error), response[1].(func() error), castError(response[2])
}

func (mDAL *MockDAL) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	addCall(mDAL, "SelectUserByEmail", email)
	response := getResponse(mDAL, "SelectUserByEmail")
	return response[0].(*models.UserProfile), castError(response[1])
}

func (mDAL *MockDAL) InsertUser(userIn *pkg.UserIn) error {
	addCall(mDAL, "InsertUser", userIn)
	response := getResponse(mDAL, "InsertUser")
	return castError(response[0])
}

func (mDAL *MockDAL) DeleteUserByPubId(userPubId *uuid.UUID) error {
	addCall(mDAL, "DeleteUserByPubId", userPubId)
	response := getResponse(mDAL, "DeleteUserByPubId")
	return castError(response[0])
}

func (mDAL *MockDAL) Migrate(dbName *string) (errVal error) {
	addCall(mDAL, "Migrate", dbName)
	response := getResponse(mDAL, "Migrate")
	return castError(response[0])
}

func (mDAL *MockDAL) InsertChurch(churchIn *pkg.ChurchIn) error {
	addCall(mDAL, "InsertChurch", churchIn)
	response := getResponse(mDAL, "InsertChurch")
	return castError(response[0])
}

func (mDAL *MockDAL) SelectChurchByEmail(email string) (church *models.ChurchProfile, errVal error) {
	addCall(mDAL, "SelectChurchByEmail", email)
	response := getResponse(mDAL, "SelectChurchByEmail")
	return response[0].(*models.ChurchProfile), castError(response[1])
}

func (mDAL *MockDAL) MakeLeader(userPubId *uuid.UUID, churchPubId *uuid.UUID) error {
	addCall(mDAL, "MakeLeader", userPubId, churchPubId)
	response := getResponse(mDAL, "MakeLeader")
	return castError(response[0])
}

func (mDAL *MockDAL) UpdateChurchActivationStatus(churchPubId *uuid.UUID, isActive bool) error {
	addCall(mDAL, "UpdateChurchActivationStatus", churchPubId, isActive)
	response := getResponse(mDAL, "UpdateChurchActivationStatus")
	return castError(response[0])
}
