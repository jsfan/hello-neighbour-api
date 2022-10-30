package models

type UserProfile struct {
	PubId        string  `json:"pub_id"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password_hash"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	DateOfBirth  *string `json:"date_of_birth"`
	Gender       string  `json:"gender"`
	Description  *string `json:"description"`
	ChurchUUID   string  `json:"church_uuid"`
	Role         string  `json:"role"`
	Active       bool    `json:"active"`
}
