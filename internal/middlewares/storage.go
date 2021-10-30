package middlewares

import (
	"context"
	"github.com/google/logger"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/endpoints"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"net/http"
)

func GetStorageMiddleWare(masterStore interfaces.DataInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check database connection first
			newStore := masterStore.Clone()
			if newStore == nil {
				logger.Error("Database connection unavailable.")
				endpoints.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			}
			ctx := context.WithValue(r.Context(), config.MasterStore, newStore)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
