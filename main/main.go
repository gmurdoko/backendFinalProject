package main

import (
	"finalproject/config"
	"finalproject/main/master"
)

func main() {
	db := config.EnvConn()
	useActivityLog := config.AuthSwitch()
	router := config.CreateRouter()
	master.Init(router, db, useActivityLog)
	config.RunServer(router)
}
