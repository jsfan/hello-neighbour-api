/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type User struct {
	Email string `json:"email"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Gender string `json:"gender,omitempty"`

	Description string `json:"description,omitempty"`

	Church string `json:"church,omitempty"`

	Role string `json:"role,omitempty"`

	DateOfBirth string `json:"date_of_birth"`

	Password string `json:"password"`

	Uuid string `json:"uuid"`

	Contact []ContactMethod `json:"contact,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"email":         obj.Email,
		"first_name":    obj.FirstName,
		"last_name":     obj.LastName,
		"date_of_birth": obj.DateOfBirth,
		"password":      obj.Password,
		"uuid":          obj.Uuid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Contact {
		if err := AssertContactMethodRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}
