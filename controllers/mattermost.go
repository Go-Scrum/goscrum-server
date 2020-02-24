package controllers

import (
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"
	"strings"

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

func (m *MattermostController) GetParticipants(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	workspaceId, err := util.GetStringKey(req.PathParameters, "workspaceId")
	if err != nil {
		return util.ServerError(err)
	}

	channelId, err := util.GetStringKey(req.PathParameters, "channelId")
	if err != nil {
		return util.ServerError(err)
	}

	participants, err := m.service.GetParticipants(workspaceId, channelId)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(participants)
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

func (m *MattermostController) GetWorkspaceByBot(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	reqToken, err := util.GetStringKey(req.Headers, "Authorization")
	if err != nil {
		return util.ServerError(err)
	}

	splitToken := strings.Split(strings.ToLower(reqToken), "bearer")
	if len(splitToken) != 2 {
		return util.ClientError(http.StatusUnauthorized)
	}

	token := strings.TrimSpace(splitToken[1])

	workspace, err := m.service.GetWorkspaceByToken(token)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(workspace)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}

func (m *MattermostController) GetParticipantQuestion(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// TODO Validate with personal token
	//reqToken, err := util.GetStringKey(req.Headers, "Authorization")
	//if err != nil {
	//	return util.ServerError(err)
	//}
	//
	//splitToken := strings.Split(strings.ToLower(reqToken), "bearer")
	//if len(splitToken) != 2 {
	//	return util.ClientError(http.StatusUnauthorized)
	//}

	projectId, err := util.GetStringKey(req.PathParameters, "projectId")
	if err != nil {
		return util.ServerError(err)
	}

	participantId, err := util.GetStringKey(req.PathParameters, "participantId")
	if err != nil {
		return util.ServerError(err)
	}

	//token := strings.TrimSpace(splitToken[1])

	workspace, err := m.service.GetParticipantQuestion(projectId, participantId)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(workspace)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}
