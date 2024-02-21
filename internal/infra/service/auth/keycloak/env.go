package keycloak

import (
	"os"
)

func GetAdminUsername() string {
	value := os.Getenv("IAM_ADMIN_USERNAME")
	if value == "" {
		value = "user"
	}
	return value
}

func GetAdminPassword() string {
	value := os.Getenv("IAM_ADMIN_PASSWORD")
	if value == "" {
		value = "bitnami"
	}
	return value
}

func GetRelmName() string {
	value := os.Getenv("IAM_RELM_NAME")
	if value == "" {
		value = "aimerzarashi"
	}
	return value
}

func GetHost() string {
	value := os.Getenv("IAM_SERVICE_HOST")
	if value == "" {
		value = "http://localhost"
	}
	return value
}

func GetClientId() string {
	value := os.Getenv("IAM_CLIENT_ID")
	if value == "" {
		value = "shop"
	}
	return value
}

func GetClientSecret() string {
	value := os.Getenv("IAM_CLIENT_SECRET")
	if value == "" {
		value = "sKiC4rmyogJ7nrqlzlTd3AKNsTfsMpeA"
	}
	return value
}
