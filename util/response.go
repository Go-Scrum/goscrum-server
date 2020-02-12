package util

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/json-iterator/go"
)

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	return ResponseError(http.StatusInternalServerError, err.Error())
}

func ResponseError(status int, data string) (events.APIGatewayProxyResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	body, _ := json.MarshalToString(map[string]string{"message": data})
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache",
		},
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func ClientError(status int) (events.APIGatewayProxyResponse, error) {
	return ResponseError(status, http.StatusText(status))
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func Success(resp string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       resp,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache",
		},
	}, nil
}

func Redirect(url string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusTemporaryRedirect,
		Body:       "",
		Headers: map[string]string{
			"Location":      url,
			"Authorization": "",
		},
	}, nil
}
