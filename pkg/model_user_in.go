/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/jsfan/hello-neighbour/internal/utils/crypto"
	"reflect"
)

type UserIn struct {
	Email string `json:"email"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Gender string `json:"gender,omitempty"`

	Description string `json:"description,omitempty"`

	Church string `json:"church,omitempty"`

	Role string `json:"role,omitempty"`

	DateOfBirth string `json:"date_of_birth"`

	Password string `json:"password"`
}

// UnmarshalJSON overrides default Unmarshal method to verify JSON fields
func (userIn *UserIn) UnmarshalJSON(data []byte) error {
	type UserIn2 UserIn
	var userIn2 UserIn2
	if err := json.Unmarshal(data, &userIn2); err != nil {
		return err
	}

	value := reflect.ValueOf(userIn2)
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		fieldValue := value.Field(i).Interface()
		if fieldValue == "" {
			return fmt.Errorf("Missing or empty field '%+v' for UserIn", fieldName)
		} else if fieldName == "Password" {
			password, err := crypto.GeneratePasswordHash(userIn2.Password)
			if err != nil {
				return err
			}
			userIn2.Password = string(password)
		}
	}

	*userIn = UserIn(userIn2)
	return nil
}
