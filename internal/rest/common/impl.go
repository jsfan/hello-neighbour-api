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

// ImplResponse response defines an error code with the associated body
type ImplResponse struct {
	Code    int
	Headers map[string][]string
	Body    interface{}
}
