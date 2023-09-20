package lambda

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/samber/lo"
	"strings"
)

func Authenticate(request events.APIGatewayV2HTTPRequest) error {
	apiKey := lo.PickBy(request.Headers, func(key string, value string) bool {
		return strings.ToLower(key) == "x-octopussubscriptionlistener-apikey"
	})

	entries := lo.Entries(apiKey)

	if len(entries) == 1 {
		err := handlers.IsAuthenticated(entries[0].Value)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("did not find the ApiKey header")
}
