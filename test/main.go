package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/TomBowyerResearchProject/common/response"
	"github.com/stretchr/testify/assert"
)

const (
	letterBytes          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomUsernameLength = 20
)

func TestRequest(
	t *testing.T,
	ts *httptest.Server,
	method,
	path string,
	body io.Reader,
) (
	*http.Response, map[string]interface{}, []response.Message,
) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)

		return nil, nil, nil
	}

	return CompleteTestRequest(t, req)
}

func CompleteTestRequest(t *testing.T, r *http.Request) (*http.Response, map[string]interface{}, []response.Message) {
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)

		return nil, nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)

		return nil, nil, nil
	}
	defer resp.Body.Close()

	var responseObj response.Response

	err = json.Unmarshal(respBody, &responseObj)
	if err != nil {
		return resp, nil, responseObj.Message
	}

	resultMap, ok := responseObj.Result.(map[string]interface{})

	if !ok {
		return resp, nil, nil
	}

	return resp, resultMap, responseObj.Message
}

func CreateNewUser(t *testing.T, url string) string {
	randomUsername := RandString(randomUsernameLength)
	fmt.Printf("Creating user %s", randomUsername)
	requestBody := strings.NewReader(
		fmt.Sprintf(
			"{\"username\": \"%s\", \"name\": \"imtom\", \"password\": \"test123\", \"secret\": \"qutCreate\" }",
			randomUsername,
		),
	)

	req, err := http.NewRequest("POST", url, requestBody)
	assert.Nil(t, err)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var responseObj response.Response

	err = json.Unmarshal(respBody, &responseObj)
	if err != nil {
		log.Fatal(err.Error())
	}

	resultMap, ok := responseObj.Result.(map[string]interface{})
	if !ok {
		log.Fatal("Failed to parse")
	}

	return fmt.Sprintf("Bearer %s", resultMap["token"].(string))
}

func RandString(n int) string {
	rand.Seed(time.Now().Unix())

	b := make([]byte, n)
	for i := range b {
		// nolint:gosec
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
