package session

import "crypto/rsa"

type jwt struct {
	rawJWT         []byte
	sessionDetails map[string]string
}

var signingKey *rsa.PrivateKey
var session *jwt

func SetSigningKey(jwtKey *rsa.PrivateKey) {
	signingKey = jwtKey
}

func (userSession *jwt) Validate() {
}

func (userSesssion *jwt) Build() {

}
