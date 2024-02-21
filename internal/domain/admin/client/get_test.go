package client_test

import (
	"keycloak-golang/internal/domain/admin/client"
	"keycloak-golang/internal/domain/admin/token"
	"testing"
)

func TestGet(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		clientId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{clientId: "shop"},
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

			// When
			id, err := client.Get(token, tt.args.clientId)

			if !tt.wantErr {
				if err != nil {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				}

				if id == "" {
					t.Errorf("Get() id = %v, wantErr %v", id, tt.wantErr)
				}
			}
		})
	}
}
