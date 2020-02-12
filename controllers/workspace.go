package controllers

import (
	"net/http"

	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
)

type WorkspaceController struct {
	service services.WorkspaceService
}

func NewWorkspaceController(service services.WorkspaceService) WorkspaceController {
	return WorkspaceController{service: service}
}

func (a *WorkspaceController) Create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	workspace := models.Workspace{}

	err := json.Unmarshal([]byte(req.Body), &workspace)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	err = a.service.Create(workspace)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	return util.Success("")
}
