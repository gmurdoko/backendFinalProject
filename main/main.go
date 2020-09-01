package main

import "finalproject/config"

func main() {
	db := config.EnvConn()
	router := config.CreateRouter()
	masters.Init(router, db)
	config.RunServer(router)
}
