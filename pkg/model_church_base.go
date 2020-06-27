/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

type ChurchBase struct {

	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`
}
