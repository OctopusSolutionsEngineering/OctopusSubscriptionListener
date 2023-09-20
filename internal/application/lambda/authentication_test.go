package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

func TestHeaderParsing(t *testing.T) {
	os.Setenv("APIKEY", "MyKey")
	err := Authenticate(events.APIGatewayV2HTTPRequest{
		Headers: map[string]string{
			"X-OctopusSubscriptionListener-ApiKey": "MyKey",
		},
	})

	if err != nil {
		t.Fatalf("RAuthenication should not have failed")
	}
}

func TestHeaderParsingFail(t *testing.T) {
	os.Setenv("APIKEY", "MyKey1")
	err := Authenticate(events.APIGatewayV2HTTPRequest{
		Headers: map[string]string{
			"X-OctopusSubscriptionListener-ApiKey": "MyKey",
		},
	})

	if err == nil {
		t.Fatalf("RAuthenication should have failed")
	}
}

func TestHeaderParsingFailEmpty(t *testing.T) {
	os.Setenv("APIKEY", "")
	err := Authenticate(events.APIGatewayV2HTTPRequest{
		Headers: map[string]string{
			"X-OctopusSubscriptionListener-ApiKey": "MyKey",
		},
	})

	if err == nil {
		t.Fatalf("RAuthenication should have failed")
	}
}
