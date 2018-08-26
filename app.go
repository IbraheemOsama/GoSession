package main

import (
	"GoSession/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router, service services.IRegisterRouterService) {
	service.InitAndRegisterService(router)
}

func initServices() []services.IRegisterRouterService {
	return []services.IRegisterRouterService{services.MoviesService{}}
}

func main() {

	router := mux.NewRouter()
	var services = initServices()
	for _, service := range services {
		registerRoutes(router, service)
	}

	if err := http.ListenAndServe(":3567", router); err != nil {
		log.Fatal(err)
	}
}
