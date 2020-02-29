package main

import (
	"goscrum/server/controllers"
	"goscrum/server/gateway"
	"goscrum/server/services"
)

func main() {

	gitlabService := services.NewGitlabPService()
	controller := controllers.NewGitlabController(gitlabService)

	router := gateway.NewAPIRouter()

	router.Get("/gitlab/{projectId}/commits", controller.Commits)
	router.Post("/gitlab/issues", controller.Issues)
	router.Get("/gitlab/users", controller.Users)
	router.Get("/gitlab/projects", controller.Projects)
	router.Get("/gitlab/{userId}/events", controller.UserEvents)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
