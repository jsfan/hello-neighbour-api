/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type ContactMethod struct {
	Label string `json:"label"`

	ContactDetail string `json:"contact_detail"`

	User string `json:"user"`

	Uuid string `json:"uuid"`
}

// AssertContactMethodRequired checks if the required fields are not zero-ed
func AssertContactMethodRequired(obj ContactMethod) error {
	elements := map[string]interface{}{
		"label":          obj.Label,
		"contact_detail": obj.ContactDetail,
		"user":           obj.User,
		"uuid":           obj.Uuid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseContactMethodRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ContactMethod (e.g. [][]ContactMethod), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseContactMethodRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aContactMethod, ok := obj.(ContactMethod)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertContactMethodRequired(aContactMethod)
	})
}
