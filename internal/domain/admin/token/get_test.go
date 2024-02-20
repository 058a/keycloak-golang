package token_test

import (
	"keycloak-golang/internal/domain/admin/token"
	"testing"
)

func TestGet(t *testing.T) {
	token, err := token.Get()
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
