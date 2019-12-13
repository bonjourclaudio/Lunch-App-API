package auth

import (
	"github.com/claudioontheweb/lunch-app/config"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth struct {
	Config *oauth2.Config
}

func NewOAuth(conf *config.GoogleConfig) *OAuth {
	oauth := OAuth{
		Config: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
		},
	}
	return &oauth
}

type AuthResponse struct {
	User	models.User
	Token	string
}