package octopus

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"net/url"
	"os"
	"strconv"
)

func LoggingEnabled(spaceId string, projectId string) (bool, error) {
	url, err := url.Parse(os.Getenv("OCTOPUS_URL"))

	if err != nil {
		return false, err
	}

	client, err := client.NewClient(nil, url, os.Getenv("OCTOPUS_APIKEY"), spaceId)

	if err != nil {
		return false, err
	}

	project, err := client.Projects.GetByID(projectId)

	if err != nil {
		return false, err
	}

	vars, err := client.Variables.GetByName(project.VariableSetID, "DemoSpaceCreator.Monitoring.Disabled", &variables.VariableScope{})

	if err != nil {
		return false, err
	}

	if len(vars) == 0 {
		return true, nil
	}

	disabled, err := strconv.ParseBool(vars[0].Value)

	return err != nil || !disabled, nil
}
