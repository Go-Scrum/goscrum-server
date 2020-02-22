package controllers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"
	"net/http"
)

type ProjectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) ProjectController {
	return ProjectController{service: service}
}

func (p *ProjectController) Save(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	project := models.Project{}

	err := json.Unmarshal([]byte(req.Body), &project)
	if err != nil {
		return util.ResponseError(http.StatusBadRequest, err.Error())
	}

	project, err = p.service.Save(project)
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

	fmt.Println("Get all")
	projects, err := p.service.GetAll()
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

func (p *ProjectController) GetbyID(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	id, err := util.GetStringKey(req.PathParameters, "id")
	if err != nil {
		return util.ServerError(err)
	}

	project, err := p.service.GetByID(id)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	resp, err := json.Marshal(project)
	if err != nil {
		return util.ResponseError(http.StatusInternalServerError, err.Error())
	}

	return util.Success(string(resp))
}
