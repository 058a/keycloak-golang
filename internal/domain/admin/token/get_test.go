package token_test

import (
	"keycloak-golang/internal/domain/admin/token"
	"testing"
)

func TestSignIn(t *testing.T) {
	token, err := token.SignIn()
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
