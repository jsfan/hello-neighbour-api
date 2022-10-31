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

type UserPublic struct {
	Email string `json:"email"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Gender string `json:"gender,omitempty"`

	Description string `json:"description,omitempty"`

	Church string `json:"church,omitempty"`

	Role string `json:"role,omitempty"`

	Uuid string `json:"uuid"`

	Contact []ContactMethod `json:"contact,omitempty"`
}

// AssertUserPublicRequired checks if the required fields are not zero-ed
func AssertUserPublicRequired(obj UserPublic) error {
	elements := map[string]interface{}{
		"email":      obj.Email,
		"first_name": obj.FirstName,
		"last_name":  obj.LastName,
		"uuid":       obj.Uuid,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	for _, el := range obj.Contact {
		if err := AssertContactMethodRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseUserPublicRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserPublic (e.g. [][]UserPublic), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserPublicRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserPublic, ok := obj.(UserPublic)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertUserPublicRequired(aUserPublic)
	})
}
