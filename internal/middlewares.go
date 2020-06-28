/*
 * Middlewares for the HTTP server (e.g. authorisation)
 */
package internal

import (
	"context"
	"encoding/json"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
	"net/http"
	"strings"
)

const authHeader = "Authorization"

func sendUnauthorizedResponse(w http.ResponseWriter) {
	resp := pkg.ErrorResponse{
		Code:    401,
		Message: "Unauthenticated",
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	_, err = w.Write(respJson)
	if err != nil {
		log.Printf("[ERROR] Could not write response for unauthenticated request: %+v", err)
	}

}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken, ok := r.Header[authHeader]
		if !ok {
			sendUnauthorizedResponse(w)
			return
		}
		// validate JWT
		// TODO: Go through all authorisation headers instead?
		tokenDetails := strings.SplitN(bearerToken[0], " ", 2)
		if strings.ToLower(tokenDetails[0]) != "bearer" {
			sendUnauthorizedResponse(w)
			return
		}
		if err := session.GetJWTWrapper().Validate(tokenDetails[1]); err != nil {
			sendUnauthorizedResponse(w)
			return
		}
		// augment request context with user session
		ctx := context.WithValue(r.Context(), config.SessionKey, "<session details to go here>")
		// add storage connection to context
		conn, err := storage.GetConnection()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("[ERROR] Database connection unavailable.")
			return
		}
		ctx = context.WithValue(ctx, config.DatabaseConnection, conn)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
