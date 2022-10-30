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

type ChurchPublicAllOf struct {
	Uuid string `json:"uuid"`
}

// AssertChurchPublicAllOfRequired checks if the required fields are not zero-ed
func AssertChurchPublicAllOfRequired(obj ChurchPublicAllOf) error {
	elements := map[string]interface{}{
		"uuid": obj.Uuid,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseChurchPublicAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ChurchPublicAllOf (e.g. [][]ChurchPublicAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseChurchPublicAllOfRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aChurchPublicAllOf, ok := obj.(ChurchPublicAllOf)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertChurchPublicAllOfRequired(aChurchPublicAllOf)
	})
}