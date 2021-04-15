package response

import (
	"encoding/json"
	"net/http"

	"github.com/TomBowyerResearchProject/common/logger"
)

type Response struct {
	Result  interface{} `json:"result"`
	Message []Message   `json:"message"`
}

type Message struct {
	Message string `json:"message"`
	Target  string `json:"target,omitempty"`
}

func ResultResponseJSON(w http.ResponseWriter, status int, result interface{}) {
	response := Response{
		Result: result,
	}

	responseJSON(w, status, response)
}

func MessageResponseJSON(w http.ResponseWriter, status int, message Message) {
	response := Response{
		Message: []Message{message},
	}

	responseJSON(w, status, response)
}

func responseJSON(w http.ResponseWriter, status int, response interface{}) {
	payload, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Sending response %v", response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(status)
	_, err = w.Write(payload)

	if err != nil {
		logger.Error(err)
	}
}
