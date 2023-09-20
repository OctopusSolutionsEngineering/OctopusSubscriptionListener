package web

import (
	"encoding/json"
	events2 "github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/events"
	"github.com/OctopusSolutionsEngineering/OctopusSubscriptionListener/internal/domain/handlers"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func HandleRequest(w http.ResponseWriter, request *http.Request) {
	if strings.ToLower(request.Method) != "post" {
		w.WriteHeader(405)
		return
	}

	body, err := io.ReadAll(request.Body)

	if err != nil {
		zap.L().Error(err.Error())
		w.WriteHeader(500)
		return
	}

	bodyString := string(body)

	zap.L().Debug(bodyString)

	subscriptionEvent := events2.SubscriptionEvent{}
	err = json.Unmarshal(body, &subscriptionEvent)

	if err != nil {
		zap.L().Error(err.Error())
		w.WriteHeader(400)
		return
	}

	err = handlers.PostToSlack(subscriptionEvent)

	if err != nil {
		zap.L().Error(err.Error())
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
