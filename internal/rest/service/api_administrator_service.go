/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"github.com/jsfan/hello-neighbour-api/internal/rest"
	"github.com/jsfan/hello-neighbour-api/internal/rest/common"
	"github.com/jsfan/hello-neighbour-api/internal/rest/model"
)

// AdministratorApiService is a service that implements the logic for the AdministratorApiServicer
// This service should implement the business logic for every endpoint for the AdministratorApi API.
// Include any external packages or services that will be required by this service.
type AdministratorApiService struct {
}

// NewAdministratorApiService creates a default api service
func NewAdministratorApiService() rest.AdministratorApiServicer {
	return &AdministratorApiService{}
}

// GetQuestions - Retrieve all questions
func (s *AdministratorApiService) GetQuestions(ctx context.Context) (common.ImplResponse, error) {
	// TODO - update GetQuestions with the required logic for this service method.
	// Add api_administrator_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Question{}) or use other options such as http.Ok ...
	//return Response(200, Question{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(403, ErrorResponse{}), nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("GetQuestions method not implemented")
}

// GetUsers - Retrieve all users
func (s *AdministratorApiService) GetUsers(ctx context.Context) (common.ImplResponse, error) {
	// TODO - update GetUsers with the required logic for this service method.
	// Add api_administrator_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []User{}) or use other options such as http.Ok ...
	//return Response(200, []User{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(403, ErrorResponse{}), nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("GetUsers method not implemented")
}

// UpdateChurchActivate - Activate or deactivate church
func (s *AdministratorApiService) UpdateChurchActivate(ctx context.Context, churchUUIDStr string, isActive model.InlineObject) (common.ImplResponse, error) {
	userSession := ctx.Value(config.SessionKey).(*config.UserSession)
	if userSession.Role != "admin" {
		return common.Response(http.StatusForbidden, map[string]interface{}{"message": "You cannot change a church's activation status.", "code": http.StatusForbidden}), nil
	}

	store := ctx.Value(config.MasterStore).(interfaces.DataInterface)

	churchUUID, err := uuid.Parse(churchUUIDStr)
	if err != nil {
		return common.Response(http.StatusBadRequest, map[string]interface{}{"message": "Invalid church UUID.", "code": http.StatusBadRequest}), nil
	}

	if err = store.ActivateChurch(ctx, &churchUUID, isActive.IsActive); err != nil {
		glog.Errorf("Could not change activation status of church %s: %+v", churchUUID.String(), err)
		return common.Response(http.StatusInternalServerError, map[string]interface{}{"message": "", "code": http.StatusInternalServerError}), nil
	}
	glog.Infof("Church %s has their activation status updated.", churchUUID.String())

	return common.Response(http.StatusNoContent, nil), nil
}