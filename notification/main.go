package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Notification struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Link     string `json:"link"`
}

func SendEvent(url string, notif Notification) error {
	bodyBytes, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
