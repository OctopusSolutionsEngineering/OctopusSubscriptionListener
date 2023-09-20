package application

import (
	"encoding/json"
	events2 "github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/events"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/handlers"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if Authenticate(request) != nil {
		return LambdaResponse(401)
	}

	zap.L().Debug(request.Body)

	subscriptionEvent := events2.SubscriptionEvent{}
	err := json.Unmarshal([]byte(request.Body), &subscriptionEvent)

	if err != nil {
		return LambdaResponse(400)
	}

	err = handlers.PostToSlack(subscriptionEvent)

	if err != nil {
		return LambdaResponse(500)
	}

	return LambdaResponse(200)
}
