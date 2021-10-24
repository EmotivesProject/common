package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	Like    = "like"
	Comment = "comment"
	Message = "message"
)

type Notification struct {
	ID         int       `json:"id,omitempty"`
	Username   string    `json:"username"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Message    string    `json:"message"`
	Link       string    `json:"link"`
	PostID     *int      `json:"post_id,omitempty"`
	UsernameTo *string   `json:"username_to,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	Seen       bool      `json:"seen"`
}

func SendEvent(url, authSecret string, notif Notification) (int, error) {
	bodyBytes, err := json.Marshal(notif)
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authSecret)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, resp.Body.Close()
}

func SendDelete(url, authSecret string) (int, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authSecret)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return resp.StatusCode, resp.Body.Close()
}
