package main

import (
	lambda2 "github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/application/lambda"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/logger"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	logger.BuildLogger()
	lambda.Start(lambda2.HandleRequest)
}
