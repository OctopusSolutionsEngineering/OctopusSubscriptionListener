package lambda

import (
	"encoding/json"
	events2 "github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/events"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/handlers"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
	"strings"
)

func HandleRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if strings.ToLower(request.RequestContext.HTTP.Method) != "post" {
		return LambdaResponse(405)
	}

	if Authenticate(request) != nil {
		return LambdaResponse(401)
	}

	zap.L().Debug(request.Body)

	subscriptionEvent := events2.SubscriptionEvent{}
	err := json.Unmarshal([]byte(request.Body), &subscriptionEvent)

	if err != nil {
		zap.L().Error(err.Error())
		return LambdaResponseCustom(400, "Failed to decode JSON body")
	}

	err = handlers.PostToSlack(subscriptionEvent)

	if err != nil {
		zap.L().Error(err.Error())
		return LambdaResponseCustom(500, "Failed to call the Slack handler, or failed to query Octopus API")
	}

	return LambdaResponse(200)
}
