package routers

import (
	"paypal-starter-backend2/m/models"
	"paypal-starter-backend2/m/services"

	"github.com/gorilla/mux"
)

var routesDBsubscriptions = []models.Route{
	{Path: "/subscription/p", Function: services.PostSubscription, Method: "POST", Mw: ""},
	{Path: "/subscriptions/g", Function: services.GetSubscriptions, Method: "GET", Mw: ""},
	{Path: "/subscription/d", Function: services.DeleteSubscription, Method: "DELETE", Mw: ""},
}

func SetRoutesDBsubscriptions(router *mux.Router) *mux.Router {
	for _, r := range routesDBsubscriptions {

		if r.Mw == "" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
		} else if r.Mw == "auth" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			//router.HandleFunc(r.Path, auth.RequireTokenAuthentication(r.Function)).Methods(r.Method)
		}

	}

	return router
}

func GetRoutesDBsubscriptions() []models.Route {
	return routesDBsubscriptions
}
