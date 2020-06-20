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
	ChurchId     int64
	Role         string
	Active       bool
}
