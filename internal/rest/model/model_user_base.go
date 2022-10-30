/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"github.com/jsfan/hello-neighbour-api/internal/rest/common"
)

type UserBase struct {
	Email string `json:"email"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Gender string `json:"gender,omitempty"`

	Description string `json:"description,omitempty"`

	Church string `json:"church,omitempty"`

	Role string `json:"role,omitempty"`
}

// AssertUserBaseRequired checks if the required fields are not zero-ed
func AssertUserBaseRequired(obj UserBase) error {
	elements := map[string]interface{}{
		"email":      obj.Email,
		"first_name": obj.FirstName,
		"last_name":  obj.LastName,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserBaseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserBase (e.g. [][]UserBase), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserBaseRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserBase, ok := obj.(UserBase)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertUserBaseRequired(aUserBase)
	})
}