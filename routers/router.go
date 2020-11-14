package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetRoutesDBsubscriptions(router)
	router = SetRoutesPaypal(router)

	return router
}
