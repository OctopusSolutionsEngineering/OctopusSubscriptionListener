package application

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func LambdaResponse(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
