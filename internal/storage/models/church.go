package models

type ChurchProfile struct {
	PubId                 string `json:"pub_id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Address               string `json:"address"`
	Website               string `json:"website"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	GroupSize             string `json:"group_size"`
	SameGender            bool   `json:"same_gender"`
	MinAge                string `json:"min_age"`
	MemberBasicInfoUpdate bool   `json:"member_basic_info_update"`
	Active                bool   `json:"active"`
}
