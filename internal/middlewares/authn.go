package middlewares

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"github.com/jsfan/hello-neighbour-api/internal/rest/common"
	"github.com/jsfan/hello-neighbour-api/internal/session"
	"github.com/jsfan/hello-neighbour-api/internal/utils"
)

func AuthzMiddleware(masterStore interfaces.DataInterface) func(handlerFunc http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			authCtx := ctx.Value(common.AuthContextKey).(*common.AuthContext)
			if authCtx == nil {
				common.EncodeJSONResponse(map[string]interface{}{"error": "Internal Server Error", "code": 500}, func(i int) *int { return &i }(http.StatusInternalServerError), nil, w)
				return
			}
			var userSession *config.UserSession
			// check if HTTP Basic or Bearer auth
			if authCtx.UserCredentials != nil {
				// check user credentials
				creds := authCtx.UserCredentials
				var authenticated bool
				userSession, authenticated = utils.CheckBasicAuth(r.Context(), masterStore.Clone(), creds.Username, creds.Password)
				if !authenticated {
					common.EncodeJSONResponse(map[string]string{"error": "Login credentials invalid."}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
					return
				}
			} else {
				// check claims
				ourJWT := session.NewJWT()
				if err := ourJWT.Validate(authCtx.ParsedJWT); err != nil {
					glog.Infof("Could not validate JWT: %+v", err)
					common.EncodeJSONResponse(map[string]string{"error": "JWT claims invalid."}, func(i int) *int { return &i }(http.StatusUnauthorized), nil, w)
					return
				}
				userSession = ourJWT.SessionDetails
			}
			ctx = context.WithValue(ctx, config.SessionKey, userSession)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
