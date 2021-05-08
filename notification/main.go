package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Title     string             `bson:"title" json:"title"`
	Message   string             `bson:"message" json:"message"`
	Link      string             `bson:"link" json:"link"`
	PostID    *int               `bson:"post_id,omitempty" json:"post_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Seen      bool               `bson:"seen" json:"seen"`
}

func SendEvent(url, authSecret string, notif Notification) error {
	bodyBytes, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authSecret)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
