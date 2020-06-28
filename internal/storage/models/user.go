package models

type UserProfile struct {
	PubId        string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	DateOfBirth  string
	Gender       string
	Description  string
	ChurchUUID   string
	Role         string
	Active       bool
}
