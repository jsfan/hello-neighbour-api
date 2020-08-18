package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(pwdHash []byte, pwd []byte) (passwordOk bool) {
	err := bcrypt.CompareHashAndPassword(pwdHash, pwd)
	if err == nil {
		return true
	}
	return false
}
