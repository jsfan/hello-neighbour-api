/*
 * COVID-19 Global Church Hack 'Hello Neighbour'
 *
 * This is the API for COVID-19 Global Church Hack 'Hello Neighbour' project
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

type Question struct {
	Question string `json:"question,omitempty"`

	Church string `json:"church,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}
