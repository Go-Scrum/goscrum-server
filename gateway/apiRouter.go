package gateway

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type APIRouter struct {
	tree map[string]ResourceMap
}

var errHandleNotFound = errors.New("handler not found")

type HandlerAPIFunc func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type Resource struct {
	handler HandlerAPIFunc
}

type ResourceMap map[string]Resource

func NewAPIRouter() APIRouter {
	return APIRouter{tree: map[string]ResourceMap{}}
}

func (r APIRouter) Get(path string, handler HandlerAPIFunc) {
	if resource, ok := r.tree[path]; ok {
		resource[http.MethodGet] = Resource{handler: handler}
	} else {
		r.tree[path] = ResourceMap{http.MethodGet: Resource{handler: handler}}
	}
}

func (r APIRouter) Post(path string, handler HandlerAPIFunc) {
	if resource, ok := r.tree[path]; ok {
		resource[http.MethodPost] = Resource{handler: handler}
	} else {
		r.tree[path] = ResourceMap{http.MethodPost: Resource{handler: handler}}
	}
}

func (r APIRouter) Put(path string, handler HandlerAPIFunc, roles []string) {
	if resource, ok := r.tree[path]; ok {
		resource[http.MethodPut] = Resource{handler: handler}
	} else {
		r.tree[path] = ResourceMap{http.MethodPut: Resource{handler: handler}}
	}
}

func (r APIRouter) Delete(path string, handler HandlerAPIFunc, roles []string) {
	if resource, ok := r.tree[path]; ok {
		resource[http.MethodDelete] = Resource{handler: handler}
	} else {
		r.tree[path] = ResourceMap{http.MethodDelete: Resource{handler: handler}}
	}
}

func (r APIRouter) GetResource(path, method string) (*Resource, error) {
	if resourceMap, ok := r.tree[path]; ok {
		if resource, ok := resourceMap[method]; ok {
			return &resource, nil
		}
	}
	return nil, errHandleNotFound
}
