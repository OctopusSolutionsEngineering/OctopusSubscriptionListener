package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/events"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/slack"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/interface/octopus"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

func PostToSlack(event events.SubscriptionEvent) error {

	projectId, foundProject := lo.Find(event.Payload.Event.RelatedDocumentIds, func(item string) bool {
		return strings.HasPrefix(item, "Projects-")
	})

	if !foundProject {
		return nil
	}

	serverTask, foundServerTask := lo.Find(event.Payload.Event.RelatedDocumentIds, func(item string) bool {
		return strings.HasPrefix(item, "ServerTasks-")
	})

	if !foundServerTask {
		return nil
	}

	spaceName, err := octopus.GetSpaceId(event.Payload.Event.SpaceId)

	if err != nil {
		zap.L().Error(err.Error())
		// Fallback to display the space ID
		spaceName = event.Payload.Event.SpaceId
	}

	enabled, err := octopus.LoggingEnabled(event.Payload.Event.SpaceId, projectId)

	if err != nil {
		zap.L().Error(err.Error())
		// Assume we want to log the event
		enabled = true
	}

	if !enabled {
		return nil
	}

	body := slack.SlackMessage{
		Channel:   os.Getenv("SLACK_CHANNEL"),
		Username:  "Demo Space Creator",
		IconUrl:   "https://octopus.com/content/resources/favicon.png",
		LinkNames: "true",
		Attachments: []slack.SlackMessageAttachments{
			{
				MrkDwnIn: []string{"pretext", "text"},
				Pretext:  "",
				Text:     spaceName + ": " + event.Payload.Event.Message + "\n<" + os.Getenv("OCTOPUS_URL") + "/app#/" + event.Payload.Event.SpaceId + "/tasks/" + serverTask + "|Task log>",
				Color:    "danger",
			},
		},
	}

	bodyJson, err := json.Marshal(body)

	if !enabled {
		return nil
	}

	requestURL := fmt.Sprintf(os.Getenv("SLACK_URL"))
	res, err := http.Post(requestURL, "application/json", bytes.NewReader(bodyJson))

	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.New("slack API call returned " + fmt.Sprint(res.StatusCode))
	}

	return nil
}
