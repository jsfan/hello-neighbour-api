// Code generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import "github.com/jsfan/hello-neighbour-api/internal/rest/common"

type ErrorResponse struct {
	Code int32 `json:"code"`

	Message string `json:"message"`
}

// AssertErrorResponseRequired checks if the required fields are not zero-ed
func AssertErrorResponseRequired(obj ErrorResponse) error {
	elements := map[string]interface{}{
		"code":    obj.Code,
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseErrorResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ErrorResponse (e.g. [][]ErrorResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseErrorResponseRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aErrorResponse, ok := obj.(ErrorResponse)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertErrorResponseRequired(aErrorResponse)
	})
}
