package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"keycloak-golang/internal/domain/admin/token"
	"keycloak-golang/internal/infra/service/auth/keycloak"
	"net/http"
)

type GetResponseBody struct {
	Id       string `json:"id"`
	ClientId string `json:"clientId"`
}

func Get(token *token.Token, cliendId string) (string, error) {
	httpClient := &http.Client{}

	url := keycloak.GetHost() + "/admin/realms/" + keycloak.GetRelmName() + "/clients"
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received status code %d", res.StatusCode)
	}

	// レスポンスボディを取得
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// JSONデータを構造体にマッピングする
	var clients []GetResponseBody
	err = json.Unmarshal(body, &clients)
	if err != nil {
		return "", err
	}

	// マッピングされた構造体を使って何かをする
	for _, client := range clients {
		if client.ClientId == cliendId {
			return client.Id, nil
		}
	}

	return "", errors.New("client not found")
}
