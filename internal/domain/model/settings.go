package model

type KeycloakSettings struct {
	BaseURL string
	Realm   string
}

type Settings struct {
	Keycloak KeycloakSettings
}

func NewSettings() *Settings {
	return &Settings{}
}
