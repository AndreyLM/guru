package v1

import (
	"github.com/andreylm/guru/pkg/routes"
	"github.com/andreylm/guru/pkg/server/v1/handlers"
)

func getRootRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{Name: "AddUser", Method: "POST", Pattern: "/user/create", HandlerFunc: handlers.AddUser},
		routes.Route{Name: "GetUser", Method: "POST", Pattern: "/user/get", HandlerFunc: handlers.GetUser},
		routes.Route{Name: "AddDeposit", Method: "POST", Pattern: "/user/deposit", HandlerFunc: handlers.AddDeposit},
		routes.Route{Name: "Transaction", Method: "POST", Pattern: "/transaction", HandlerFunc: handlers.Transaction},
	}
}
