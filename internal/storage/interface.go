package storage

import "github.com/jsfan/hello-neighbour/pkg"

type DBInteraction interface {
	GetUserByEmail(username string)
	UserRegister(userIn *pkg.UserIn)
}
