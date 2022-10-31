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

type ChurchIn struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	GroupSize int32 `json:"group_size"`

	SameGender bool `json:"same_gender"`

	MinAge int32 `json:"min_age"`

	MemberBasicInfoUpdate bool `json:"member_basic_info_update"`
}

// AssertChurchInRequired checks if the required fields are not zero-ed
func AssertChurchInRequired(obj ChurchIn) error {
	elements := map[string]interface{}{
		"name":                     obj.Name,
		"description":              obj.Description,
		"address":                  obj.Address,
		"group_size":               obj.GroupSize,
		"same_gender":              obj.SameGender,
		"min_age":                  obj.MinAge,
		"member_basic_info_update": obj.MemberBasicInfoUpdate,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseChurchInRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ChurchIn (e.g. [][]ChurchIn), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseChurchInRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aChurchIn, ok := obj.(ChurchIn)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertChurchInRequired(aChurchIn)
	})
}
