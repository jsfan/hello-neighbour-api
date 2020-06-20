package session

type jwt struct {
	rawJWT         []byte
	sessionDetails map[string]string
}

var session *jwt

func (userSession *jwt) Validate() {

}

func (userSesssion *jwt) Build() {

}
