package response

import (
	"encoding/json"
	"net/http"

	"github.com/TomBowyerResearchProject/common/logger"
)

const (
	healthzResponse = "Health OK"
)

type Response struct {
	Result  interface{} `json:"result"`
	Message []Message   `json:"message"`
}

type Message struct {
	Message string `json:"message"`
	Target  string `json:"target,omitempty"`
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	MessageResponseJSON(w, false, http.StatusOK, Message{Message: healthzResponse})
}

func ResultResponseJSON(w http.ResponseWriter, cache bool, status int, result interface{}) {
	response := Response{
		Result: result,
	}

	responseJSON(w, cache, status, response)
}

func MessageResponseJSON(w http.ResponseWriter, cache bool, status int, message Message) {
	response := Response{
		Message: []Message{message},
	}

	responseJSON(w, cache, status, response)
}

func responseJSON(w http.ResponseWriter, cache bool, status int, response interface{}) {
	payload, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)

		return
	}

	logger.Infof("Sending response %v", string(payload))

	w.Header().Set("Content-Type", "application/json")

	if !cache {
		w.Header().Set("Cache-Control", "no-cache")
	}

	w.WriteHeader(status)
	_, err = w.Write(payload)

	if err != nil {
		logger.Error(err)
	}
}
