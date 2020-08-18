package session

import (
	"crypto/rsa"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

const issuer = "Hello Neighbour Team"
const subject = "Hello Neighbour API"

type jwtWrapper struct {
	rawJWT         string
	SessionDetails *UserSession
}

// JWT lifespan (i.e. time until it expires)
var jwtLifespan, _ = time.ParseDuration("24h")

// Key used to sign JWT
var signingKey *rsa.PrivateKey

func SetSigningKey(jwtKey *rsa.PrivateKey) {
	signingKey = jwtKey
}

func NewJWT() *jwtWrapper {
	return &jwtWrapper{
		rawJWT:         "",
		SessionDetails: nil,
	}
}

func (jwtWrapper *jwtWrapper) Validate(rawJWT string) error {
	tok, err := jwt.ParseSigned(rawJWT)
	if err != nil {
		return errors.Wrap(err, "could not parse JWT")
	}

	expectedBase := jwt.Expected{
		Issuer:  issuer,
		Subject: subject,
		Time:    time.Now(),
	}
	ourClaims := UserSession{}
	if err := tok.Claims(&signingKey.PublicKey, &expectedBase, &ourClaims); err != nil {
		return errors.Wrap(err, "could not validate claims")
	}

	jwtWrapper.SessionDetails = &ourClaims
	return nil
}

func (jwtWrapper *jwtWrapper) Build(sessionClaims *UserSession) error {
	if signingKey == nil {
		return errors.New("no signing key loaded")
	}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: signingKey}, nil)
	if err != nil {
		return errors.Wrap(err, "could not create JWT signer")
	}
	claims := jwt.Claims{
		Issuer:   issuer,
		Subject:  subject,
		Expiry:   jwt.NewNumericDate(time.Now().Add(jwtLifespan)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}
	jwtWrapper.rawJWT, err = jwt.Signed(signer).Claims(claims).Claims(sessionClaims).CompactSerialize()
	return errors.Wrap(err, "could not build JWT")
}

func (jwtWrapper *jwtWrapper) GetRaw() string {
	return jwtWrapper.rawJWT
}
