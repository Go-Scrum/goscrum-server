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
	service := services.NewMattermostService(workspaceService)

	controller := controllers.NewMattermostController(service)
	router := gateway.NewAPIRouter()

	router.Get("/mattermost/{workspaceId}/channels", controller.GetChannels)
	router.Get("/mattermost/{workspaceId}/teams", controller.GetTeams)
	router.Get("/mattermost/{workspaceId}/{channelId}/participants", controller.GetParticipants)
	router.Get("/mattermost/bot", controller.GetWorkspaceByBot)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
