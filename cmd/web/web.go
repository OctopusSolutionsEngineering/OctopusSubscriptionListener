package main

import (
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/application/web"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	logger.BuildLogger()
	http.HandleFunc("/", web.HandleRequest)
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		zap.L().Error(err.Error())
		os.Exit(1)
	}
}
