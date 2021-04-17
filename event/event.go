package event

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/TomBowyerResearchProject/common/logger"
)

var (
	eventConfig EventConfig
)

type EventConfig struct {
	EventURL            string
	AuthorizationHeader string
}

type Event struct {
	Username      string      `json:"username"`
	CustomerEvent string      `json:"customer_event"`
	Data          interface{} `json:"event_data"`
}

type EventData struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func Init(cfg EventConfig) {
	eventConfig = cfg
}

func SendEvent(event Event) {
	requestBody, err := json.Marshal(event)
	if err != nil {
		logger.Error(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", eventConfig.EventURL, bytes.NewReader(requestBody))
	if err != nil {
		logger.Error(err)
	}
	req.Header.Add("Authorization", eventConfig.AuthorizationHeader)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	logger.Info("Sent event to metrics")
}
