package models

type Auth struct {
	Model
	URL          string `json:"url"`
	AuthType     int    `json:"auth_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}
