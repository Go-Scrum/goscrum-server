package main

import (
	"goscrum/server/controllers"
	"goscrum/server/db"
	"goscrum/server/gateway"
	"goscrum/server/services"
)

func main() {
	db := db.DbClient(true)

	defer db.Close()

	workspaceService := services.NewWorkspaceService(db)
	service := services.NewBotService(workspaceService)

	controller := controllers.NewMattermostController(service)
	router := gateway.NewAPIRouter()

	router.Get("/mattermost/{workspaceId}/channels", controller.GetChannels)
	router.Get("/mattermost/{workspaceId}/teams", controller.GetTeams)
	router.Get("/mattermost/project", controller.GetTeams)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
