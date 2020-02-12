package gateway

import (
	"net/http"

	"goscrum/server/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Gateway struct {
}

func NewGateway() Gateway {
	return Gateway{}
}

func (g Gateway) StartAPI(r APIRouter) {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		resource, err := r.GetResource(request.Resource, request.HTTPMethod)
		if err != nil {
			return util.ClientError(http.StatusMethodNotAllowed)
		}
		var response events.APIGatewayProxyResponse
		response, err = resource.handler(request)
		return response, err
	})
}
