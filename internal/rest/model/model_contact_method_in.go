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

type ContactMethodIn struct {
	Label string `json:"label"`

	ContactDetail string `json:"contact_detail"`

	User string `json:"user"`
}

// AssertContactMethodInRequired checks if the required fields are not zero-ed
func AssertContactMethodInRequired(obj ContactMethodIn) error {
	elements := map[string]interface{}{
		"label":          obj.Label,
		"contact_detail": obj.ContactDetail,
		"user":           obj.User,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseContactMethodInRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ContactMethodIn (e.g. [][]ContactMethodIn), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseContactMethodInRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aContactMethodIn, ok := obj.(ContactMethodIn)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertContactMethodInRequired(aContactMethodIn)
	})
}
