package config

import (
	"finalproject/utils/getEnv"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateRouter creating new router for app.goCreateRouter
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

//RunServer running the server
func RunServer(router *mux.Router) {
	appHost := getEnv.ViperGetEnv("APP_HOST", "localhost")
	appPort := getEnv.ViperGetEnv("APP_PORT", "8080")
	hostListen := fmt.Sprintf("%v:%v", appHost, appPort)
	log.Printf("Ready to listen on %v", hostListen)
	err := http.ListenAndServe(hostListen, router)
	if err != nil {
		log.Fatal(err)
	}
}
