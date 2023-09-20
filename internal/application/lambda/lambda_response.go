package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func LambdaResponse(status int) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func LambdaResponseCustom(status int, statusText string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       statusText,
	}, nil
}
