/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type ChurchBase struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`
}

// AssertChurchBaseRequired checks if the required fields are not zero-ed
func AssertChurchBaseRequired(obj ChurchBase) error {
	elements := map[string]interface{}{
		"name":        obj.Name,
		"description": obj.Description,
		"address":     obj.Address,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseChurchBaseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ChurchBase (e.g. [][]ChurchBase), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseChurchBaseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aChurchBase, ok := obj.(ChurchBase)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertChurchBaseRequired(aChurchBase)
	})
}
