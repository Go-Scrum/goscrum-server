package controllers

import (
	"fmt"
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

func (a *WorkspaceController) Save(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	workspace := models.Workspace{}
	if claimPayload, ok := req.RequestContext.Authorizer["claims"]; ok {
		claims := claimPayload.(map[string]interface{})
		if claims != nil {
			workspace.UserEmail = fmt.Sprintf("%s", claims["email"])

			err := json.Unmarshal([]byte(req.Body), &workspace)
			if err != nil {
				return util.ResponseError(http.StatusBadRequest, err.Error())
			}

			newWorkspace, err := a.service.Save(workspace)
			if err != nil {
				return util.ResponseError(http.StatusInternalServerError, err.Error())
			}

			result, err := json.MarshalToString(newWorkspace)
			if err != nil {
				return util.ResponseError(http.StatusBadRequest, err.Error())
			}

			return util.Success(result)
		}
	}
	return util.ClientError(http.StatusUnauthorized)
}

func (a *WorkspaceController) Get(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if claimPayload, ok := req.RequestContext.Authorizer["claims"]; ok {
		claims := claimPayload.(map[string]interface{})
		if claims != nil {
			email := fmt.Sprintf("%s", claims["email"])
			workspace, err := a.service.GetWorkspaceByUserEmail(email)
			if err != nil {
				return util.ResponseError(http.StatusInternalServerError, err.Error())
			}

			result, err := json.MarshalToString(workspace)
			if err != nil {
				return util.ResponseError(http.StatusBadRequest, err.Error())
			}

			return util.Success(result)
		}
	}
	return util.ClientError(http.StatusUnauthorized)
}
