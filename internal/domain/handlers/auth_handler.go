package handlers

import (
	"errors"
	"os"
)

func IsAuthenticated(apiKey string) error {
	if apiKey == os.Getenv("APIKEY") {
		return nil
	}

	return errors.New("supplied API Key does not match")
}
