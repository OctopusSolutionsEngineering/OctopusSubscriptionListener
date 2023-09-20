package octopus

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"strconv"
)

func LoggingEnabled(spaceId string, projectId string) (bool, error) {
	client, err := getClient(spaceId)

	if err != nil {
		return false, err
	}

	vars, err := client.Variables.GetByName(projectId, "DemoSpaceCreator.Monitoring.Disabled", &variables.VariableScope{})

	if err != nil {
		return false, err
	}

	if len(vars) == 0 {
		return true, nil
	}

	disabled, err := strconv.ParseBool(vars[0].Value)

	return err != nil || !disabled, nil
}

func GetSpaceId(spaceId string) (string, error) {
	client, err := getClient(spaceId)

	if err != nil {
		return "", err
	}

	space, err := client.Spaces.GetByID(spaceId)

	if err != nil {
		return "", err
	}

	return space.Name, nil
}
