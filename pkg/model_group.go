/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

type Group struct {
	Uuid string `json:"uuid"`

	Created string `json:"created"`

	Users []UserPublic `json:"users"`

	Questions []Question `json:"questions,omitempty"`
}
