// Code generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package common

import (
	"reflect"
)

// Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{
		Code:    code,
		Headers: nil,
		Body:    body,
	}
}

// ResponseWithHeaders return a ImplResponse struct filled, including headers
func ResponseWithHeaders(code int, headers map[string][]string, body interface{}) ImplResponse {
	return ImplResponse{
		Code:    code,
		Headers: headers,
		Body:    body,
	}
}

// IsZeroValue checks if the val is the zero-ed value.
func IsZeroValue(val interface{}) bool {
	return val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

// AssertRecurseInterfaceRequired recursively checks each struct in a slice against the callback.
// This method traverse nested slices in a preorder fashion.
func AssertRecurseInterfaceRequired(obj interface{}, callback func(interface{}) error) error {
	return AssertRecurseValueRequired(reflect.ValueOf(obj), callback)
}

// AssertRecurseValueRequired checks each struct in the nested slice against the callback.
// This method traverse nested slices in a preorder fashion.
func AssertRecurseValueRequired(value reflect.Value, callback func(interface{}) error) error {
	switch value.Kind() {
	// If it is a struct we check using callback
	case reflect.Struct:
		if err := callback(value.Interface()); err != nil {
			return err
		}

	// If it is a slice we continue recursion
	case reflect.Slice:
		for i := 0; i < value.Len(); i += 1 {
			if err := AssertRecurseValueRequired(value.Index(i), callback); err != nil {
				return err
			}
		}
	}
	return nil
}
