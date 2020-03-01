package controllers

import (
	"fmt"
	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davecgh/go-spew/spew"
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

	teamId, err := util.GetStringKey(req.PathParameters, "teamId")
	if err != nil {
		return util.ServerError(err)
	}

	channels, err := m.service.GetPublicChannelsForTeam(workspaceId, teamId)
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

func (m *MattermostController) BotUserAction(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
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
	fmt.Println(string(req.Body))

	return util.Success(`{
	  "update": {
		"message": "Updated!"
	  },
	  "ephemeral_text": "You updated the post!"
	}`)
}

func (m *MattermostController) GetQuestionDetails(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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

	//token := strings.TrimSpace(splitToken[1])

	questionId, err := util.GetStringKey(req.PathParameters, "questionId")
	if err != nil {
		return util.ServerError(err)
	}

	question, err := m.service.GetQuestionDetails(questionId)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(question)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}

func (m *MattermostController) AddUserActivity(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	userActivity := models.UserActivity{}

	err := json.Unmarshal([]byte(req.Body), &userActivity)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	err = m.service.AddUserActivity(userActivity)
	if err != nil {
		return util.ServerError(err)
	}

	return util.Success("")
}

func (m *MattermostController) UserInteraction(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	//
	//token := strings.TrimSpace(splitToken[1])

	userId, err := util.GetStringKey(req.PathParameters, "userId")
	if err != nil {
		return util.ServerError(err)
	}

	//workspace, err := m.service.GetWorkspaceByToken(token)
	//if err != nil {
	//	return util.ServerError(err)
	//}

	message := models.Message{}

	err = json.Unmarshal([]byte(req.Body), &message)
	spew.Dump(message)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	newMessage, err := m.service.UserInteraction(userId, message)
	spew.Dump(newMessage)
	spew.Dump(err)
	if err != nil {
		return util.ServerError(err)
	}

	result, err := json.MarshalToString(newMessage)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	return util.Success(result)
}
