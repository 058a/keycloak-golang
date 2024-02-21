package token

import (
	"encoding/json"
	"fmt"
	"io"
	"keycloak-golang/internal/infra/service/auth/keycloak"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func SignIn() (*Token, error) {
	httpClient := &http.Client{}
	values := url.Values{}
	values.Add("client_id", "admin-cli")
	values.Add("grant_type", "password")
	values.Add("username", keycloak.GetAdminUsername())
	values.Add("password", keycloak.GetAdminPassword())

	url := keycloak.GetHost() + "/realms/master/protocol/openid-connect/token"
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	token := &Token{}
	json.Unmarshal(resBodyByte, &token)

	return token, nil
}
