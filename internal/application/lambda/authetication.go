package lambda

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/handlers"
	"github.com/aws/aws-lambda-go/events"
)

func Authenticate(request events.APIGatewayV2HTTPRequest) error {
	if apiKey, ok := request.Headers["X-OctopusSubscriptionListener-ApiKey"]; ok {
		err := handlers.IsAuthenticated(apiKey)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("did not find the ApiKey header")
}
