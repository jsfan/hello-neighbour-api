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

type ChurchPublic struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	Uuid string `json:"uuid"`
}

// AssertChurchPublicRequired checks if the required fields are not zero-ed
func AssertChurchPublicRequired(obj ChurchPublic) error {
	elements := map[string]interface{}{
		"name":        obj.Name,
		"description": obj.Description,
		"address":     obj.Address,
		"uuid":        obj.Uuid,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseChurchPublicRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ChurchPublic (e.g. [][]ChurchPublic), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseChurchPublicRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aChurchPublic, ok := obj.(ChurchPublic)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertChurchPublicRequired(aChurchPublic)
	})
}
