/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsfan/hello-neighbour-api/internal/rest/common"
	"github.com/jsfan/hello-neighbour-api/internal/rest/model"

	"github.com/gorilla/mux"
)

// MemberApiController binds http requests to an api service and writes the service results to the http response
type MemberApiController struct {
	service      MemberApiServicer
	errorHandler common.ErrorHandler
}

// MemberApiOption for how the controller is set up.
type MemberApiOption func(*MemberApiController)

// WithMemberApiErrorHandler inject ErrorHandler into controller
func WithMemberApiErrorHandler(h common.ErrorHandler) MemberApiOption {
	return func(c *MemberApiController) {
		c.errorHandler = h
	}
}

// NewMemberApiController creates a default api controller
func NewMemberApiController(s MemberApiServicer, opts ...MemberApiOption) common.Router {
	controller := &MemberApiController{
		service:      s,
		errorHandler: common.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MemberApiController
func (c *MemberApiController) Routes() common.Routes {
	return common.Routes{
		{
			"AcceptInvite",
			strings.ToUpper("Patch"),
			"/v0/register/{userUUID}",
			c.AcceptInvite,
			true,
		},
		{
			"AddChurch",
			strings.ToUpper("Post"),
			"/v0/church",
			c.AddChurch,
			true,
		},
		{
			"AddContactMethod",
			strings.ToUpper("Post"),
			"/v0/user/{userUUID}/contactmethod",
			c.AddContactMethod,
			true,
		},
		{
			"DeleteContactMethod",
			strings.ToUpper("Delete"),
			"/v0/user/{userUUID}/contactmethod/{methodUUID}",
			c.DeleteContactMethod,
			true,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/v0/user/{userUUID}",
			c.DeleteUser,
			true,
		},
		{
			"EditUser",
			strings.ToUpper("Put"),
			"/v0/user/{userUUID}",
			c.EditUser,
			true,
		},
		{
			"GetChurches",
			strings.ToUpper("Get"),
			"/v0/church",
			c.GetChurches,
			true,
		},
		{
			"GetMatchGroup",
			strings.ToUpper("Get"),
			"/v0/user/{userUUID}/matchGroup",
			c.GetMatchGroup,
			true,
		},
		{
			"GetMessages",
			strings.ToUpper("Get"),
			"/v0/user/{userUUID}/matchgroup/{groupUUID}/bulletin",
			c.GetMessages,
			true,
		},
		{
			"GetUser",
			strings.ToUpper("Get"),
			"/v0/user/{userUUID}",
			c.GetUser,
			true,
		},
		{
			"LoginUser",
			strings.ToUpper("Get"),
			"/v0/login",
			c.LoginUser,
			true,
		},
		{
			"SendMessage",
			strings.ToUpper("Post"),
			"/v0/user/{userUUID}/matchgroup/{groupUUID}/bulletin",
			c.SendMessage,
			true,
		},
		{
			"UpdateContactMethod",
			strings.ToUpper("Put"),
			"/v0/user/{userUUID}/contactmethod/{methodUUID}",
			c.UpdateContactMethod,
			true,
		},
		{
			"UserProfile",
			strings.ToUpper("Get"),
			"/v0/profile",
			c.UserProfile,
			true,
		},
	}
}

// AcceptInvite - Accept invite
func (c *MemberApiController) AcceptInvite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	bodyParam := model.UserIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertUserInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AcceptInvite(r.Context(), userUUIDParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// AddChurch - Request new church
func (c *MemberApiController) AddChurch(w http.ResponseWriter, r *http.Request) {
	bodyParam := model.ChurchIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertChurchInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddChurch(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// AddContactMethod - Add a contact method to a user profile
func (c *MemberApiController) AddContactMethod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	bodyParam := model.ContactMethodIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertContactMethodInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddContactMethod(r.Context(), userUUIDParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// DeleteContactMethod - Delete a contact method from a user profile
func (c *MemberApiController) DeleteContactMethod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	methodUUIDParam := params["methodUUID"]

	result, err := c.service.DeleteContactMethod(r.Context(), userUUIDParam, methodUUIDParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// DeleteUser - Delete user
func (c *MemberApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	result, err := c.service.DeleteUser(r.Context(), userUUIDParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// EditUser - Update user
func (c *MemberApiController) EditUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	bodyParam := model.UserIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertUserInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.EditUser(r.Context(), userUUIDParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetChurches - Retrieve all signed up churches
func (c *MemberApiController) GetChurches(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetChurches(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetMatchGroup - Retrieve current match group
func (c *MemberApiController) GetMatchGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	result, err := c.service.GetMatchGroup(r.Context(), userUUIDParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetMessages - Retrieve all messages
func (c *MemberApiController) GetMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	groupUUIDParam := params["groupUUID"]

	result, err := c.service.GetMessages(r.Context(), userUUIDParam, groupUUIDParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetUser - Retrieve user details
func (c *MemberApiController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	result, err := c.service.GetUser(r.Context(), userUUIDParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// LoginUser - Login user
func (c *MemberApiController) LoginUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.LoginUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// SendMessage - Send message
func (c *MemberApiController) SendMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	groupUUIDParam := params["groupUUID"]

	bodyParam := model.MessageIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertMessageInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SendMessage(r.Context(), userUUIDParam, groupUUIDParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// UpdateContactMethod - Update a contact method for a user
func (c *MemberApiController) UpdateContactMethod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUIDParam := params["userUUID"]

	methodUUIDParam := params["methodUUID"]

	bodyParam := model.ContactMethodIn{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := model.AssertContactMethodInRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateContactMethod(r.Context(), userUUIDParam, methodUUIDParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// UserProfile - Logged in user's UUID and church UUID
func (c *MemberApiController) UserProfile(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserProfile(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	common.EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
