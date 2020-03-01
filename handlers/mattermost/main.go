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
	projectService := services.NewProjectService(db)
	userActivityService := services.NewUserActivityService(db)
	participantService := services.NewParticipantService(db)
	service := services.NewMattermostService(workspaceService, projectService, userActivityService, participantService)

	controller := controllers.NewMattermostController(service)
	router := gateway.NewAPIRouter()

	router.Get("/mattermost/{workspaceId}/teams", controller.GetTeams)
	router.Get("/mattermost/{workspaceId}/{teamId}/channels", controller.GetChannels)
	router.Get("/mattermost/{workspaceId}/channel/{channelId}/participants", controller.GetParticipants)
	router.Get("/mattermost/bot", controller.GetWorkspaceByBot)
	router.Post("/mattermost/bot/user/{userId}/action", controller.BotUserAction)
	router.Get("/mattermost/bot/user/{userId}/action", controller.BotUserAction)
	router.Get("/mattermost/bot/question/{questionId}", controller.GetQuestionDetails)
	router.Post("/mattermost/bot/user/activity", controller.AddUserActivity)
	router.Post("/mattermost/bot/user/{userId}/message", controller.UserInteraction)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
