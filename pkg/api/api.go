package api

import (
	"strings"

	"github.com/gkkkb/piston/middleware"
	
	"github.com/julienschmidt/httprouter"
)

// API contains informations needed for an API
type API struct {
	Endpoint  string
	Action    string
	Method    string
	Authority Authority
	Handle    HandleWithError
}

// Authority represents authority of users
type Authority int

const (
	// Anonymous represents users not logged in
	Anonymous Authority = iota
	// User represents normal users
	User
	// Admin represents trusted users (admins)
	Admin
)

// StartAPIs starts API handlers
func StartAPIs(router *httprouter.Router, apis []API) {
	for _, api := range apis {
		var (
			action string
			handle HandleWithError
		)

		if strings.HasPrefix(api.Endpoint, "/_internal") {
			action, handle = newInternalHandle(api.Action, api.Authority, api.Handle)
		} else {
			action, handle = newHandle(api.Action, api.Authority, api.Handle)
		}

		router.Handle(api.Method, api.Endpoint, middleware.MonitorHTTP(action, handle))
	}
}
