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

	service := services.NewProjectService(db)

	controller := controllers.NewProjectController(service)
	router := gateway.NewAPIRouter()

	router.Post("/projects", controller.Save)
	router.Get("/projects", controller.GetAll)
	router.Get("/projects/{id}", controller.GetbyID)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
