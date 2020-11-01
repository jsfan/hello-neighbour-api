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
	"reflect"
)

type ChurchIn struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	GroupSize int32 `json:"group_size"`

	SameGender bool `json:"same_gender"`

	MinAge int32 `json:"min_age"`

	MemberBasicInfoUpdate bool `json:"member_basic_info_update"`
}

// UnmarshalJSON overrides default Unmarshal method to verify JSON fields
func (churchIn *ChurchIn) UnmarshalJSON(data []byte) error {
	type ChurchIn2 ChurchIn
	var churchIn2 ChurchIn2
	if err := json.Unmarshal(data, &churchIn2); err != nil {
		return err
	}

	value := reflect.ValueOf(churchIn2)
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		fieldValue := value.Field(i).Interface()
		if fieldValue == "" {
			return fmt.Errorf("Missing or empty field '%+v' for ChurchIn", fieldName)
		}
	}

	*churchIn = ChurchIn(churchIn2)
	return nil
}
