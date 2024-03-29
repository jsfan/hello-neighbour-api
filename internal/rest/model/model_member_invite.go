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

type MemberInvite struct {
	Email string `json:"email"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	DateOfBirth string `json:"date_of_birth,omitempty"`
}

// AssertMemberInviteRequired checks if the required fields are not zero-ed
func AssertMemberInviteRequired(obj MemberInvite) error {
	elements := map[string]interface{}{
		"email": obj.Email,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseMemberInviteRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MemberInvite (e.g. [][]MemberInvite), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMemberInviteRequired(objSlice interface{}) error {
	return common.AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMemberInvite, ok := obj.(MemberInvite)
		if !ok {
			return common.ErrTypeAssertionError
		}
		return AssertMemberInviteRequired(aMemberInvite)
	})
}
