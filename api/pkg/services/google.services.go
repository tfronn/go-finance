package services

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOAuth2Config() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		RedirectURL:  "http://localhost:4000/api/v1/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return conf

}
