/*
 * Middlewares for the HTTP server (e.g. authorisation)
 */
package internal

import (
	"context"
	"encoding/json"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
	"net/http"
)

const authHeader = "Authorization"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Header[authHeader]
		if !ok {
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
