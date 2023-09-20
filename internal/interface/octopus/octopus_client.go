package octopus

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"net/url"
	"os"
)

func getClient(spaceId string) (*client.Client, error) {
	parsedUrl, err := url.Parse(os.Getenv("OCTOPUS_URL"))

	if err != nil {
		return nil, err
	}

	return client.NewClient(nil, parsedUrl, os.Getenv("OCTOPUS_APIKEY"), spaceId)
}
