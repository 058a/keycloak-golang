package token_test

import (
	"testing"

	"keycloak-golang/internal/domain/account/token"
)

func TestSignIn(t *testing.T) {
	t.Parallel()

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{username: "16b329a2-d94d-4b3e-84fe-3389fd9424b5"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := token.SignIn(tt.args.username)

			// Then
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.AccessToken == "" {
				t.Error("access token is empty")
			}
			if got.RefreshToken == "" {
				t.Error("refresh token is empty")
			}
		})
	}
}
