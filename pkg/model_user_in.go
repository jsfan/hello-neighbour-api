/*
 * COVID-19 Global Church Hack 'Hello Neighbour'
 *
 * This is the API for COVID-19 Global Church Hack 'Hello Neighbour' project
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

type UserIn struct {
	Email string `json:"email,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Description string `json:"description,omitempty"`

	Church string `json:"church,omitempty"`

	DateOfBirth string `json:"dateOfBirth,omitempty"`
}
