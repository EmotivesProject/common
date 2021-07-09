package verification

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/TomBowyerResearchProject/common/response"
)

type key string

const (
	UserID key = "username"
)

var (
	errUnauthorised    = errors.New("Not authorised")
	verificationConfig VerificationConfig
)

type VerificationUser struct {
	Username string `json:"username"`
}

type VerificationConfig struct {
	VerificationURL     string
	AuthorizationSecret string
}

func Init(cfg VerificationConfig) {
	verificationConfig = cfg
}

func VerifyJTW() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client := &http.Client{}

			header := r.Header.Get("Authorization")

			req, err := http.NewRequest("GET", verificationConfig.VerificationURL, nil)
			if err != nil {
				response.MessageResponseJSON(w, false, http.StatusInternalServerError, response.Message{
					Message: err.Error(),
				})

				return
			}

			req.Header.Add("Authorization", header)
			resp, err := client.Do(req)
			if err != nil {
				response.MessageResponseJSON(w, false, http.StatusInternalServerError, response.Message{
					Message: err.Error(),
				})

				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					response.MessageResponseJSON(w, false, http.StatusInternalServerError, response.Message{
						Message: err.Error(),
					})

					return
				}

				var dat response.Response
				var user VerificationUser
				_ = json.Unmarshal(bodyBytes, &dat)
				body, _ := json.Marshal(dat.Result)

				_ = json.Unmarshal(body, &user)
				ctx := context.WithValue(r.Context(), UserID, user.Username)
				next.ServeHTTP(w, r.WithContext(ctx))

				return
			}

			response.MessageResponseJSON(w, false, http.StatusUnauthorized, response.Message{
				Message: errUnauthorised.Error(),
			})
		})
	}
}

func VerifyToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == verificationConfig.AuthorizationSecret {
				next.ServeHTTP(w, r)

				return
			}

			response.MessageResponseJSON(w, false, http.StatusBadRequest, response.Message{
				Message: errUnauthorised.Error(),
			})
		})
	}
}
