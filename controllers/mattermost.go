package controllers

import (
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
)

type MattermostController struct {
	service services.MattermostService
}

func NewMattermostController(service services.MattermostService) MattermostController {
	return MattermostController{service: service}
}

func (m *MattermostController) GetChannels(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	workspaceId, err := util.GetStringKey(req.PathParameters, "workspaceId")
	if err != nil {
		return util.ServerError(err)
	}

	channels, err := m.service.GetAllPublicChannels(workspaceId)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(channels)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}

func (m *MattermostController) GetTeams(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	workspaceId, err := util.GetStringKey(req.PathParameters, "workspaceId")
	if err != nil {
		return util.ServerError(err)
	}

	teams, err := m.service.GetAllTeams(workspaceId)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(teams)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}
