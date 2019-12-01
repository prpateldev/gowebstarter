package routerutils

import "github.com/gorilla/mux"

import "log"

var router *mux.Router

func GetRouter() *mux.Router {
	return router
}

func SetRouter(r *mux.Router) {
	router = r
}

func Init() {
	SetRouter(mux.NewRouter())
	log.Println("Router initiated")
}
