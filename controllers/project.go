package controllers

import (
	"fmt"
	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
)

type ProjectController struct {
	projectService   services.ProjectService
	workspaceService services.WorkspaceService
}

func NewProjectController(projectService services.ProjectService, workspaceService services.WorkspaceService) ProjectController {
	return ProjectController{projectService: projectService, workspaceService: workspaceService}
}

func (p *ProjectController) Save(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	email := getEmailFromClaim(req)
	if os.Getenv("LOCAL") == "true" {
		email = "durgaprasad.budhwani@gmail.com"
	}

	if email == "" {
		return util.ClientError(http.StatusUnauthorized)
	}

	workspace, err := p.workspaceService.GetWorkspaceByUserEmail(email)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	project := models.Project{
		WorkspaceID: workspace.ID,
	}

	err = json.Unmarshal([]byte(req.Body), &project)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	project, err = p.projectService.Save(project)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	resp, err := json.Marshal(project)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	return util.Success(string(resp))
}

func (p *ProjectController) GetAll(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	email := getEmailFromClaim(req)
	if os.Getenv("LOCAL") == "true" {
		email = "durgaprasad.budhwani@gmail.com"
	}

	if email == "" {
		return util.ClientError(http.StatusUnauthorized)
	}

	workspace, err := p.workspaceService.GetWorkspaceByUserEmail(email)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("Get all")
	projects, err := p.projectService.GetAll(workspace.ID)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(projects)
	resp, err := json.Marshal(projects)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}
	return util.Success(string(resp))
}

func (p *ProjectController) GetById(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO -- check for user authentication using claims
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	id, err := util.GetStringKey(req.PathParameters, "id")
	if err != nil {
		return util.ServerError(err)
	}

	project, err := p.projectService.GetByID(id)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(project)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}

func (p *ProjectController) GetParticipantQuestion(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	id, err := util.GetStringKey(req.PathParameters, "id")
	if err != nil {
		return util.ServerError(err)
	}

	project, err := p.projectService.GetByID(id)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(project)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}
