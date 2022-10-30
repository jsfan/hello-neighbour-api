/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type UserProfile struct {
	UserUuid string `json:"user_uuid"`

	ChurchUuid string `json:"church_uuid"`
}

// AssertUserProfileRequired checks if the required fields are not zero-ed
func AssertUserProfileRequired(obj UserProfile) error {
	elements := map[string]interface{}{
		"user_uuid":   obj.UserUuid,
		"church_uuid": obj.ChurchUuid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserProfileRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserProfile (e.g. [][]UserProfile), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserProfileRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserProfile, ok := obj.(UserProfile)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserProfileRequired(aUserProfile)
	})
}
