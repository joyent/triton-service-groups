package server

import (
	"net/http"

	"github.com/joyent/triton-service-groups/groups"
	"github.com/joyent/triton-service-groups/server/router"
	"github.com/joyent/triton-service-groups/templates"
)

var templateRoutes = router.Routes{
	router.Route{
		Name:    "ListTemplates",
		Method:  http.MethodGet,
		Pattern: "/v1/tsg/templates",
		Handler: templates_v1.List,
	},
	router.Route{
		Name:    "GetTemplate",
		Method:  http.MethodGet,
		Pattern: "/v1/tsg/templates/{identifier}",
		Handler: templates_v1.Get,
	},
	router.Route{
		Name:    "CreateTemplate",
		Method:  http.MethodPost,
		Pattern: "/v1/tsg/templates",
		Handler: templates_v1.Create,
	},
	router.Route{
		Name:    "DeleteTemplate",
		Method:  http.MethodDelete,
		Pattern: "/v1/tsg/templates/{identifier}",
		Handler: templates_v1.Delete,
	},
}

var groupRoutes = router.Routes{
	router.Route{
		Name:    "GetGroup",
		Method:  http.MethodGet,
		Pattern: "/v1/tsg/{identifier}",
		Handler: groups_v1.Get,
	},
	router.Route{
		Name:    "CreateGroup",
		Method:  http.MethodPost,
		Pattern: "/v1/tsg",
		Handler: groups_v1.Create,
	},
	router.Route{
		Name:    "UpdateGroup",
		Method:  http.MethodPut,
		Pattern: "/v1/tsg/{identifier}",
		Handler: groups_v1.Update,
	},
	router.Route{
		Name:    "DeleteGroup",
		Method:  http.MethodDelete,
		Pattern: "/v1/tsg/{identifier}",
		Handler: groups_v1.Delete,
	},
	router.Route{
		Name:    "ListGroups",
		Method:  http.MethodGet,
		Pattern: "/v1/tsg",
		Handler: groups_v1.List,
	},
}

var RoutingTable = router.RouteTable{
	templateRoutes,
	groupRoutes,
}
