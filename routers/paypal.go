package routers

import (
	"paypal-starter-backend2/m/models"
	"paypal-starter-backend2/m/services"

	"github.com/gorilla/mux"
)

var routesPaypal = []models.Route{
	//{Path: "/paypal/token", Function: services.RequestToken, Method: "GET", Mw: "auth"},
	{Path: "/paypal/products", Function: services.GetProducts, Method: "GET", Mw: "auth"},
	{Path: "/paypal/plans", Function: services.GetPlans, Method: "GET", Mw: "plans"},
	{Path: "/paypal/sub/cancel", Function: services.CancelSubscription, Method: "POST", Mw: "subs"},
	{Path: "/paypal/sub/details", Function: services.GetSubscriptionDetails, Method: "GET", Mw: "subs"},
}

func SetRoutesPaypal(router *mux.Router) *mux.Router {
	for _, r := range routesPaypal {

		if r.Mw == "" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
		} else if r.Mw == "auth" {
			//router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			router.HandleFunc(r.Path, services.RequestToken(r.Function)).Methods(r.Method)
		} else if r.Mw == "plans" {
			//router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			router.HandleFunc(r.Path, services.RequestToken(r.Function)).Methods(r.Method).Queries("product_id", "{product_id}")

		} else if r.Mw == "subs" {
			//router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			router.HandleFunc(r.Path, services.RequestToken(r.Function)).Methods(r.Method).Queries("subscriptionID", "{subscriptionID}")

		}

	}

	return router
}

func GetRoutesPaypal() []models.Route {
	return routesPaypal
}
