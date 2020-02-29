package controllers

import (
	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"
	"time"
)

type GitlabController struct {
	gitlab *services.GitlabService
}

func NewGitlabController(gitlab *services.GitlabService) *GitlabController {
	return &GitlabController{gitlab: gitlab}
}

func (g *GitlabController) Issues(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	issue := models.RequestIssues{}

	err := json.Unmarshal([]byte(req.Body), &issue)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}
	issues, err := g.gitlab.Issues(issue)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	resp, err := json.Marshal(issues)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	return util.Success(string(resp))
}

func (g *GitlabController) Commits(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	id, err := util.GetStringKey(req.PathParameters, "projectId")
	if err != nil {
		return util.ServerError(err)
	}

	r := models.RequestCommits{Since: time.Now().AddDate(0, 0, -1)}
	commits, err := g.gitlab.Commits(r, id)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(commits)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}

func (g *GitlabController) Users(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	search, err := util.GetStringKey(req.QueryStringParameters, "search")
	if err != nil {
		return util.ServerError(err)
	}

	users, err := g.gitlab.Users(search)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(users)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}

func (g *GitlabController) Projects(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	search, err := util.GetStringKey(req.QueryStringParameters, "search")
	if err != nil {
		return util.ServerError(err)
	}

	users, err := g.gitlab.Projects(search)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(users)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}

func (g *GitlabController) UserEvents(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	action, err := util.GetStringKey(req.QueryStringParameters, "action")
	if err != nil {
		return util.ServerError(err)
	}

	userId, err := util.GetStringKey(req.PathParameters, "userId")
	if err != nil {
		return util.ServerError(err)
	}

	users, err := g.gitlab.UserContributions(userId, action, time.Now().AddDate(0, 0, -1))
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(users)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}
