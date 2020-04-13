/*
 * Middlewares for the HTTP server (e.g. authorisation)
 */
package internal

import (
	"encoding/json"
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
				log.Printf("Could not write response for unauthenticated request: %+v", err)
			}
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
