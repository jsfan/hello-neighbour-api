/*
 * Middlewares for the HTTP server (e.g. authorisation)
 */
package internal

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/endpoints"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/internal/utils"
	"log"
	"net/http"
	"strings"
)

const authHeader = "Authorization"

func sendUnauthorizedResponse(w http.ResponseWriter) {
	endpoints.SendErrorResponse(w, http.StatusUnauthorized, "Unauthenticated")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check database connection first
		if _, err := storage.GetConnection(); err != nil {
			log.Printf("[ERROR] Database connection unavailable.")
			endpoints.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		}

		bearerToken, ok := r.Header[authHeader]
		if !ok {
			if r.RequestURI == "/v0/user" {
				next.ServeHTTP(w, r)
			} else {
				sendUnauthorizedResponse(w)
			}
			return
		}
		// validate JWT
		// TODO: Go through all authorisation headers instead?
		tokenDetails := strings.SplitN(bearerToken[0], " ", 2)
		tokenType := strings.ToLower(tokenDetails[0])
		var userSession *session.UserSession
		switch tokenType {
		case "bearer":
			ourJwt := session.NewJWT()
			if err := ourJwt.Validate(tokenDetails[1]); err != nil {
				log.Printf("[INFO] Could not validate JWT: %+v", err)
				sendUnauthorizedResponse(w)
				return
			}
			userSession = ourJwt.SessionDetails
		case "basic":
			var authFail bool
			if userSession, authFail = utils.CheckBasicAuth(r); authFail == true {
				sendUnauthorizedResponse(w)
				return
			} else if userSession == nil {
				endpoints.SendErrorResponse(w, http.StatusBadRequest, "Malformed basic authentication")
				return
			}
		default:
			sendUnauthorizedResponse(w)
			return
		}
		// augment request context with user session
		ctx := context.WithValue(r.Context(), config.SessionKey, userSession)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
