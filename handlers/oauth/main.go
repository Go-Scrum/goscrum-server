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

	controller := controllers.NewAuthController(service)
	router := gateway.NewAPIRouter()

	router.Get("/oauth/mattermost/{workspaceId}/login", controller.MattermostLogin)
	router.Get("/oauth/mattermost/callback", controller.MattermostOauth)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
