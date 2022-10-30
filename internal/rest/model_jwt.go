/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type Jwt struct {
	Jwt string `json:"jwt"`
}

// AssertJwtRequired checks if the required fields are not zero-ed
func AssertJwtRequired(obj Jwt) error {
	elements := map[string]interface{}{
		"jwt": obj.Jwt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseJwtRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Jwt (e.g. [][]Jwt), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseJwtRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aJwt, ok := obj.(Jwt)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertJwtRequired(aJwt)
	})
}
