package middlewares

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"github.com/jsfan/hello-neighbour-api/internal/rest/common"
)

func StorageMiddleWare(masterStore interfaces.DataInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check database connection first
			newStore := masterStore.Clone()
			if newStore == nil {
				glog.Error("Database connection unavailable.")
				common.EncodeJSONResponse(`{"error": "Internal Server Error"}`, func(i int) *int { return &i }(http.StatusInternalServerError), nil, w)
			}
			ctx := context.WithValue(r.Context(), config.MasterStore, newStore)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
