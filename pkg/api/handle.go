package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gkkkb/pokedex/pkg/api/response"
	"github.com/gkkkb/pokedex/pkg/constants"
	"github.com/gkkkb/pokedex/pkg/currentuser" 
	"github.com/gkkkb/pokedex/pkg/log"
	"github.com/gkkkb/pokedex/pkg/resource"

	"github.com/julienschmidt/httprouter"
)

// HandleWithError is placeholder type while waiting for packen v1.0.4
// https://github.com/bukalapak/packen/tree/master/middleware
type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

func newHandle(action string, security Authority, handle HandleWithError) (string, HandleWithError) {
	return action, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {

		startTime := time.Now()
		rID := r.Context().Value("X-Request-ID").(string)
		ctx := resource.NewContext(r.Context(), rID, action, startTime)

		currentUser, err := currentuser.FromRequest(r)
		if err != nil {
			log.ErrLog(ctx, err, "authorization", "authorize fail")
			response.Write(w, response.BuildError([]error{response.InvalidTokenError}), response.InvalidTokenError.HTTPCode)
			return response.InvalidTokenError
		}
		
		ctx = currentuser.NewContext(ctx, currentUser)

		if !isRequestAuthorized(security, currentUser) {
			response.Write(w, response.BuildError([]error{response.UserUnauthorizedError}), response.UserUnauthorizedError.HTTPCode)
			return response.UserUnauthorizedError
		}

		r = r.WithContext(ctx)
		return handle(w, r, params)
	}
}

func newInternalHandle(action string, security Authority, handle HandleWithError) (string, HandleWithError) {
	return action, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
		headerUsername, headerPass, ok := r.BasicAuth()

		if !ok {
			response.Write(w, response.BuildError([]error{response.UserUnauthorizedError}), response.UserUnauthorizedError.HTTPCode)
			return response.UserUnauthorizedError
		}

		username := os.Getenv("POKEDEX_USERNAME")
		pass := os.Getenv("POKEDEX_PASSWORD")

		if username != headerUsername || pass != headerPass {
			response.Write(w, response.BuildError([]error{response.UserUnauthorizedError}), response.UserUnauthorizedError.HTTPCode)
			return response.UserUnauthorizedError
		}

		return handle(w, r, params)
	}
}

func isRequestAuthorized(security Authority, currentUser *currentuser.CurrentUser) bool {
	switch security {
	case Admin:
		return isRoleAllowed(currentUser.Role)
	case User:
		return isUserLoggedIn(currentUser.ID)
	default:
		return true
	}
}

func isRoleAllowed(role string) bool {
	return isInSliceString(role, []string{constants.ROLE_ADM})
}

func isUserLoggedIn(userID uint) bool {
	return userID != 0
}
