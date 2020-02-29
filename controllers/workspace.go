package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

type WorkspaceController struct {
	service services.WorkspaceService
}

func NewWorkspaceController(service services.WorkspaceService) WorkspaceController {
	return WorkspaceController{service: service}
}

func (a *WorkspaceController) Save(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	email := getEmailFromClaim(req)
	if os.Getenv("LOCAL") == "true" {
		email = "durgaprasad.budhwani@gmail.com"
	}

	if email == "" {
		return util.ClientError(http.StatusUnauthorized)
	}

	workspace := models.Workspace{}
	workspace.UserEmail = email
	workspace.URL = strings.TrimSuffix(workspace.URL, "/")

	err := json.Unmarshal([]byte(req.Body), &workspace)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	newWorkspace, err := a.service.Save(workspace)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	conf := util.GetMatterMostOAuthClient(newWorkspace)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL(newWorkspace.ID, oauth2.AccessTypeOffline)
	urlMap := map[string]string{
		"url": url,
	}
	result, err := json.MarshalToString(urlMap)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}
	return util.Success(result)
}

func getEmailFromClaim(req events.APIGatewayProxyRequest) string {
	if claimPayload, ok := req.RequestContext.Authorizer["claims"]; ok {
		claims := claimPayload.(map[string]interface{})
		if claims != nil {
			return fmt.Sprintf("%s", claims["email"])
		}
	}
	return ""
}

func (a *WorkspaceController) Get(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	email := getEmailFromClaim(req)
	if os.Getenv("LOCAL") == "true" {
		email = "durgaprasad.budhwani@gmail.com"
	}

	if email == "" {
		return util.ClientError(http.StatusUnauthorized)
	}

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
