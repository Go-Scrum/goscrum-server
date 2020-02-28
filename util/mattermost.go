package util

import (
	"fmt"
	"goscrum/server/constants"
	"goscrum/server/models"
	"os"

	"golang.org/x/oauth2"
)

func GetMatterMostOAuthClient(workspace models.Workspace) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     workspace.ClientID,
		ClientSecret: workspace.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/oauth/access_token", workspace.URL),
			AuthURL:  fmt.Sprintf("%s/oauth/authorize", workspace.URL),
		},
		RedirectURL: fmt.Sprintf("%s/oauth/mattermost/callback", os.Getenv(constants.ApiUrl)),
	}
}
