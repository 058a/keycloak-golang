package secret

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"keycloak-golang/internal/domain/admin/token"
	"keycloak-golang/internal/infra/service/auth/keycloak"
	"net/http"
)

type (
	RegenerateReqBody struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
)

func Regenerate(token *token.Token, id string) (string, error) {
	httpClient := &http.Client{}

	url := keycloak.GetHost() + "/admin/realms/" + keycloak.GetRelmName() + "/clients/" + id + "/client-secret"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
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
	var regenerateReqBody RegenerateReqBody
	err = json.Unmarshal(body, &regenerateReqBody)
	if err != nil {
		return "", err
	}

	return regenerateReqBody.Value, nil
}
