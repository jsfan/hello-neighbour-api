package common

import (
	"context"
	"net/http"
	"strings"
	"time"

	"gopkg.in/square/go-jose.v2/jwt"
)

type AuthContextLabel string

type BasicAuthCredentials struct {
	Username string
	Password string
}

type AuthContext struct {
	UserCredentials *BasicAuthCredentials
	ParsedJWT       *jwt.JSONWebToken
}

const (
	AuthContextKey           AuthContextLabel = "AuthContext"
	authHeader                                = "Authorization"
	authzUnknownError                         = "Authorization type unknown."
	basicAuthzMalformedError                  = "HTTP Basic header malformed."
	jwtExpiredError                           = "JWT expired."
	jwtParseError                             = "JWT parsing failed."
	jwtSignatureInvalidError                  = "JWT signature invalid."
	malformedAuthzError                       = "Authorization header malformed."
	noTokenError                              = "No bearer token present."
)

func AuthnMiddleware(next http.Handler, jwtKeys interface{}, authnRequired bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken, ok := r.Header[authHeader]
		if !ok {
			if !authnRequired {
				next.ServeHTTP(w, r)
			} else {
				EncodeJSONResponse(map[string]string{"error": noTokenError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
			}
			return
		}

		// validate JWT
		tokenDetails := strings.SplitN(bearerToken[0], " ", 2)

		if len(tokenDetails) < 2 {
			EncodeJSONResponse(map[string]string{"error": malformedAuthzError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
			return
		}

		tokenType := strings.ToLower(tokenDetails[0])
		token := tokenDetails[1]
		var authContext *AuthContext
		switch tokenType {
		case "bearer":
			// parse signed JWT
			parsedJWT, err := jwt.ParseSigned(token)
			if err != nil {
				EncodeJSONResponse(map[string]string{"error": jwtParseError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
				return
			}

			// check signature and extract claims
			basicClaims := &jwt.Claims{}
			if err := parsedJWT.Claims(jwtKeys, &basicClaims); err != nil {
				EncodeJSONResponse(map[string]string{"error": jwtSignatureInvalidError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
				return
			}

			// validate timestamp
			expected := jwt.Expected{
				Time: time.Now(),
			}
			if err := basicClaims.Validate(expected); err != nil {
				EncodeJSONResponse(map[string]string{"error": jwtExpiredError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
				return
			}
		case "basic":
			username, password, ok := r.BasicAuth()
			if !ok {
				EncodeJSONResponse(map[string]string{"error": basicAuthzMalformedError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
				return
			}
			authContext.UserCredentials = &BasicAuthCredentials{
				Username: username,
				Password: password,
			}
		default:
			EncodeJSONResponse(map[string]string{"error": authzUnknownError}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
			return
		}
		// augment request context with auth context
		ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}