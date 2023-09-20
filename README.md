This project creates a proxy that receives subscription events from Octopus and triggers another service in response.

# Environment Variables
* `APIKEY`: The Api Key that must be sent as the `X-OctopusSubscriptionListener-ApiKey` header.
* `OCTOPUS_URL`: The Octopus server
* `OCTOPUS_APIKEY`: The Octopus API key
* `SLACK_URL`: The Slack webhook URL
* `SLACK_CHANNEL`: The Slack channel

# Interfaces

This app can be built either as a web app accepting POST requests to http://localhost:333 or an AWS Lambda. 

The web version is built with `go build cmd/web/web.go`, and the Lambda version is built with `go build cmd/lambda/lambda.go`.

# Sample Payload

The JSON below is a sample of the webhook data sent by Octopus.

```json
{
  "EventType": "SubscriptionPayload",
  "Payload": {
    "BatchId": "9d0ecd22-4208-474d-8a0b-10c9c267f6be",
    "BatchProcessingDate": "2023-09-19T21:35:24.279+00:00",
    "Event": {
      "Category": "DeploymentFailed",
      "ChangeDetails": {
        "Differences": null,
        "DocumentContext": null
      },
      "Comments": null,
      "Details": null,
      "Id": "Events-1111977",
      "IdentityEstablishedWith": "",
      "IpAddress": null,
      "IsService": false,
      "Links": {
        "Self": "/api/events/Events-1111977"
      },
      "Message": "Deploy to Development failed for Argo CD Octopub release 0.1.178 to Development",
      "MessageHtml": "<a href='#/deployments/Deployments-12519'>Deploy to Development</a> failed for <a href='#/projects/Projects-4522'>Argo CD Octopub</a> release <a href='#/releases/Releases-11138'>0.1.178</a> to <a href='#/environments/Environments-1621'>Development</a>",
      "MessageReferences": [
        {
          "Length": 21,
          "ReferencedDocumentId": "Deployments-12519",
          "StartIndex": 0
        },
        {
          "Length": 15,
          "ReferencedDocumentId": "Projects-4522",
          "StartIndex": 33
        },
        {
          "Length": 7,
          "ReferencedDocumentId": "Releases-11138",
          "StartIndex": 57
        },
        {
          "Length": 11,
          "ReferencedDocumentId": "Environments-1621",
          "StartIndex": 68
        }
      ],
      "Occurred": "2023-09-19T21:35:08.563+00:00",
      "RelatedDocumentIds": [
        "Deployments-12519",
        "Projects-4522",
        "Releases-11138",
        "Environments-1621",
        "ServerTasks-569388",
        "Channels-5462"
      ],
      "SpaceId": "Spaces-282",
      "UserAgent": "Server",
      "UserId": "users-system",
      "Username": "system"
    },
    "EventNumberInBatch": 1,
    "ServerAuditUri": "https://mattc.octopus.app/app#/Spaces-282/configuration/audit?eventCategories=DeploymentFailed&from=2023-09-19T21%3a34%3a53.%2b00%3a00&to=2023-09-19T21%3a35%3a24.%2b00%3a00",
    "ServerUri": "https://mattc.octopus.app",
    "Subscription": {
      "EventNotificationSubscription": {
        "EmailDigestLastProcessed": null,
        "EmailDigestLastProcessedEventAutoId": null,
        "EmailFrequencyPeriod": "01:00:00",
        "EmailPriority": "Normal",
        "EmailShowDatesInTimeZoneId": "UTC",
        "EmailTeams": [],
        "Filter": {
          "DocumentTypes": [],
          "Environments": [],
          "EventAgents": [],
          "EventCategories": [
            "DeploymentFailed"
          ],
          "EventGroups": [],
          "ProjectGroups": [],
          "Projects": [],
          "Tags": [],
          "Tenants": [],
          "Users": []
        },
        "WebhookHeaderKey": null,
        "WebhookHeaderValue": null,
        "WebhookLastProcessed": "2023-09-19T21:34:53.464+00:00",
        "WebhookLastProcessedEventAutoId": 1117952,
        "WebhookTeams": [],
        "WebhookTimeout": "00:00:10",
        "WebhookURI": "https://eo1ekhz9261va71.m.pipedream.net"
      },
      "Id": "Subscriptions-61",
      "IsDisabled": false,
      "Links": {
        "Self": "/api/Spaces-282/subscriptions/Subscriptions-61"
      },
      "Name": "Deployment",
      "SpaceId": "Spaces-282",
      "Type": "Event"
    },
    "TotalEventsInBatch": 1
  },
  "Timestamp": "2023-09-19T21:35:25.215+00:00"
}
```