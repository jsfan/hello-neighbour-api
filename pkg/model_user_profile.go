/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

type UserProfile struct {
	UserUuid string `json:"user_uuid"`

	ChurchUuid string `json:"church_uuid"`
}