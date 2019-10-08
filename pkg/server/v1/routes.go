package v1

import (
	"github.com/andreylm/guru/pkg/db"
	"github.com/andreylm/guru/pkg/routes"
	"github.com/andreylm/guru/pkg/server/v1/handlers"
)

func getRootRoutes(dbStorage db.Storage) routes.Routes {
	return routes.Routes{
		routes.Route{Name: "AddUser", Method: "POST", Pattern: "/user/create", HandlerFunc: handlers.AddUser(dbStorage)},
		routes.Route{Name: "GetUser", Method: "POST", Pattern: "/user/get", HandlerFunc: handlers.GetUser(dbStorage)},
		routes.Route{Name: "AddDeposit", Method: "POST", Pattern: "/user/deposit", HandlerFunc: handlers.AddDeposit(dbStorage)},
		routes.Route{Name: "Transaction", Method: "POST", Pattern: "/transaction", HandlerFunc: handlers.Transaction(dbStorage)},
	}
}
