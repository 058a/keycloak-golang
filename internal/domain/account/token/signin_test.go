package token_test

import (
	"keycloak-golang/internal/domain/account/token"
	"testing"
)

func TestSignIn(t *testing.T) {
	username := "16b329a2-d94d-4b3e-84fe-3389fd9424b5"
	token, err := token.SignIn(username)
	if err != nil {
		t.Error(err)
	}

	if token.AccessToken == "" {
		t.Error("access token is empty")
	}

	if token.RefreshToken == "" {
		t.Error("refresh token is empty")
	}
}
