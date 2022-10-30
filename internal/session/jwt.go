package session

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/jsfan/hello-neighbour-api/internal/config"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

const issuer = "Hello Neighbour Team"
const subject = "Hello Neighbour API"

type jwtWrapper struct {
	rawJWT         string
	SessionDetails *config.UserSession
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

func (jwtWrapper *jwtWrapper) Validate(parsedJWT *jwt.JSONWebToken) error {
	expectedBase := jwt.Expected{
		Issuer:  issuer,
		Subject: subject,
		Time:    time.Now(),
	}
	ourClaims := config.UserSession{}
	if err := parsedJWT.Claims(&signingKey.PublicKey, &expectedBase, &ourClaims); err != nil {
		return err
	}

	jwtWrapper.SessionDetails = &ourClaims
	return nil
}

func (jwtWrapper *jwtWrapper) Build(sessionClaims *config.UserSession) error {
	if signingKey == nil {
		return fmt.Errorf("no signing key loaded")
	}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: signingKey}, nil)
	if err != nil {
		return fmt.Errorf("could not create JWT signer: %w", err)
	}
	claims := jwt.Claims{
		Issuer:   issuer,
		Subject:  subject,
		Expiry:   jwt.NewNumericDate(time.Now().Add(jwtLifespan)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}
	jwtWrapper.rawJWT, err = jwt.Signed(signer).Claims(claims).Claims(sessionClaims).CompactSerialize()
	return fmt.Errorf("could not build JWT: %w", err)
}

func (jwtWrapper *jwtWrapper) GetRaw() string {
	return jwtWrapper.rawJWT
}
