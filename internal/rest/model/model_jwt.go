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

type Jwt struct {
	Jwt string `json:"jwt"`
}

// AssertJwtRequired checks if the required fields are not zero-ed
func AssertJwtRequired(obj Jwt) error {
	elements := map[string]interface{}{
		"jwt": obj.Jwt,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseJwtRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Jwt (e.g. [][]Jwt), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseJwtRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aJwt, ok := obj.(Jwt)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertJwtRequired(aJwt)
	})
}