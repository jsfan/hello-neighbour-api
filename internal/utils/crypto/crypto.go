package crypto

import "golang.org/x/crypto/bcrypt"

func CheckPassword(pwdHash []byte, pwd []byte) (passwordOk bool) {
	err := bcrypt.CompareHashAndPassword(pwdHash, pwd)
	return err == nil
}

func GeneratePasswordHash(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
}
