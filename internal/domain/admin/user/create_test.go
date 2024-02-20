package user_test

import (
	"keycloak-golang/internal/domain/admin/token"
	"keycloak-golang/internal/domain/admin/user"
	"testing"
)

func TestCreate(t *testing.T) {
	token, err := token.Get()
	if err != nil {
		t.Fatal(err)
	}

	if err := user.Create(token); err != nil {
		t.Error(err)
	}
}
