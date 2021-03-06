package models

type ChurchProfile struct {
	PubId                 string
	Name                  string
	Description           string
	Address               string
	Website               string
	Email                 string
	Phone                 string
	GroupSize             string
	SameGender            bool
	MinAge                string
	MemberBasicInfoUpdate bool
	Active                bool
}
