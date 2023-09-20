package main

import (
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/application"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/logger"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	logger.BuildLogger()
	lambda.Start(application.HandleRequest)
}
