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

	projectService := services.NewProjectService(db)
	workspaceService := services.NewWorkspaceService(db)

	controller := controllers.NewProjectController(projectService, workspaceService)
	router := gateway.NewAPIRouter()

	router.Post("/projects", controller.Save)
	router.Get("/projects", controller.GetAll)
	router.Get("/projects/{id}", controller.GetById)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
