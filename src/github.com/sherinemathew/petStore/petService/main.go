package main

import (
        "fmt"
		"github.com/sherinemathew/petStore/petService/service"
		"github.com/sherinemathew/petStore/petService/dbClient"
)

var appName = "petService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
    service.StartWebServer("6004")

}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}