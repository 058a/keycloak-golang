package secret_test

import (
	"fmt"
	"keycloak-golang/internal/domain/admin/client"
	"keycloak-golang/internal/domain/admin/client/secret"
	"keycloak-golang/internal/domain/admin/token"
	"testing"
)

func TestRegenerate(t *testing.T) {
	// Setup
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			token, err := token.SignIn()
			if err != nil {
				t.Fatal(err)
			}

			id, err := client.Get(token, "shop")
			if err != nil {
				t.Fatal(err)
			}

			// When
			secret, err := secret.Regenerate(token, id)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("Regenerate() error = %v, wantErr %v", err, tt.wantErr)
				}

				if secret == "" {
					t.Errorf("Regenerate() secret = %v, wantErr %v", secret, tt.wantErr)
				}

				fmt.Printf("secret: %s\n", secret)
			}
		})
	}
}
