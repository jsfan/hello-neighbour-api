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

func addCall(mDAL MockDAL, functionName string, args ...interface{}) {
	if mDAL.Calls == nil {
		mDAL.Calls = []*CallSignature{{
			FunctionName: functionName,
			Args:         args,
		}}
	}
}

func getResponse(mDAL MockDAL, functionName string) []interface{} {
	response := mDAL.Responses[functionName][0]
	mDAL.Responses[functionName] = mDAL.Responses[functionName][1:]
	return response
}

func castError(rawError interface{}) error {
	typedError, _ := rawError.(error)
	return typedError
}

func (mDAL MockDAL) SetupDal(ctx context.Context) (commit func() error, errVal error) {
	addCall(mDAL, "SetupDAL", ctx)
	response := getResponse(mDAL, "SetupDAL")
	return response[0].(func() error), castError(response[1])
}

func (mDAL MockDAL) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	addCall(mDAL, "SelectUserByEmail", email)
	response := getResponse(mDAL, "SelectUserByEmail")
	return response[0].(*models.UserProfile), castError(response[1])
}

func (mDAL MockDAL) InsertUser(userIn *pkg.UserIn) error {
	addCall(mDAL, "InsertUser", userIn)
	response := getResponse(mDAL, "InsertUser")
	return castError(response[0])
}

func (mDAL MockDAL) DeleteUserByPubId(userPubId *uuid.UUID) error {
	addCall(mDAL, "DeleteUserByPubId", userPubId)
	response := getResponse(mDAL, "DeleteUserByPubId")
	return castError(response[0])
}

func (mDAL MockDAL) Migrate(dbName *string) (errVal error) {
	addCall(mDAL, "Migrate", dbName)
	response := getResponse(mDAL, "Migrate")
	return castError(response[0])
}
