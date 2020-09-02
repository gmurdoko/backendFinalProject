package main

import (
	"finalproject/config"
	"finalproject/main/master"
)

func main() {
	db := config.EnvConn()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
