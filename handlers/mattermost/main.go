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
	service := services.NewMattermostService(workspaceService, projectService)

	controller := controllers.NewMattermostController(service)
	router := gateway.NewAPIRouter()

	router.Get("/mattermost/{workspaceId}/teams", controller.GetTeams)
	router.Get("/mattermost/{workspaceId}/{teamId}/channels", controller.GetChannels)
	router.Get("/mattermost/{workspaceId}/{channelId}/participants", controller.GetParticipants)
	router.Get("/mattermost/bot", controller.GetWorkspaceByBot)
	router.Get("/mattermost/bot/{projectId}/{participantId}/question", controller.GetParticipantQuestion)
	router.Get("/mattermost/bot/question/{questionId}", controller.GetQuestionDetails)
	router.Post("/mattermost/bot/answer/post", controller.UpdateAnswerPostId)
	router.Post("/mattermost/bot/user/{userId}/message", controller.UserInteraction)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
