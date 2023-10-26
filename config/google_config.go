package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleLoginConfig oauth2.Config

func LoadGoogleConfig(cfg *Env) oauth2.Config {
	GoogleLoginConfig = oauth2.Config{
		RedirectURL:  cfg.Google.Redirect,
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return GoogleLoginConfig
}
