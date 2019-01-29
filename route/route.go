package route

import (
	"github.com/gkkkb/pokedex/pkg/api"
	"github.com/gkkkb/pokedex/pkg/pokedex"
)

func Route() []api.API {
	apis := []api.API{
		//{Endpoint: "/profiles", Action: "call-profiles-all", Method: "GET", Authority: api.Admin, Handle: pokedex.AllProfilesAdvanced},
		//{Endpoint: "/profiles/:profile_id", Action: "call-profile-detail", Method: "GET", Authority: api.User, Handle: pokedex.DetailProfile}
		//{Endpoint: "/_internal/autos/users/:username/status", Action: "call-user-status-by-username", Method: "GET", Authority: api.Anonymous, Handle: decepticon.UserStatus},
		//{Endpoint: "/_internal/autos/users/:username/proposals/:proposal_vehicle_type/status", Action: "call-user-capability-to-create-proposal", Method: "GET", Authority: api.Anonymous, Handle: decepticon.UserPermissionToCreateProposal},
	}

	return apis
}
