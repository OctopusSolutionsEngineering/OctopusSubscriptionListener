package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

func TestHeaderParsing(t *testing.T) {
	os.Setenv("APIKEY", "MyKey")
	err := Authenticate(events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"X-OctopusSubscriptionListener-ApiKey": "MyKey",
		},
	})

	if err != nil {
		t.Fatalf("RAuthenication should not have failed")
	}
}
