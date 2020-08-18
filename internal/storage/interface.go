package storage

type DBInteraction interface {
	GetUserByEmail(username string)
}
