/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rest

type Question struct {

	Question string `json:"question"`

	Church string `json:"church"`

	Uuid string `json:"uuid"`
}

// AssertQuestionRequired checks if the required fields are not zero-ed
func AssertQuestionRequired(obj Question) error {
	elements := map[string]interface{}{
		"question": obj.Question,
		"church": obj.Church,
		"uuid": obj.Uuid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseQuestionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Question (e.g. [][]Question), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQuestionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQuestion, ok := obj.(Question)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQuestionRequired(aQuestion)
	})
}
