package main

import (
	"fmt"

	"goscrum/server/controllers"
	"goscrum/server/db"
	"goscrum/server/gateway"
	"goscrum/server/services"
)

func main() {
	fmt.Println("Request Initiated")

	db := db.DbClient(true)

	defer db.Close()

	service := services.NewWorkspaceService(db)

	controller := controllers.NewWorkspaceController(service)
	router := gateway.NewAPIRouter()

	router.Post("/workspace", controller.Save)
	router.Get("/workspace", controller.Get)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
