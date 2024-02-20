package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keycloak-golang/internal/domain/admin/token"
	"keycloak-golang/internal/infra/service/auth/keycloak"
	"net/http"

	"github.com/google/uuid"
)

type (
	body struct {
		Username      string `json:"username"`
		Enabled       string `json:"enabled"`
		Email         string `json:"email"`
		EmailVerified string `json:"emailVerified"`
		attributes    struct {
			AccountId string `json:"accountId"`
		}
	}
)

func Create(token *token.Token) error {
	httpClient := &http.Client{}

	reqBody := body{
		Username:      uuid.NewString(),
		Enabled:       "true",
		Email:         "test" + "@example.com",
		EmailVerified: "true",
		attributes: struct {
			AccountId string `json:"accountId"`
		}{
			AccountId: uuid.NewString(),
		},
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := keycloak.GetHost() + "/admin/realms/" + keycloak.GetRelmName() + "/users"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("received status code %d", res.StatusCode)
	}

	return nil
}
