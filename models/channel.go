package models

const (
	CHANNEL_OPEN    = "O"
	CHANNEL_PRIVATE = "P"
	CHANNEL_DIRECT  = "D"
	CHANNEL_GROUP   = "G"
)

type Channel struct {
	Id          string `json:"id"`
	TeamId      string `json:"team_id"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Header      string `json:"header"`
	Purpose     string `json:"purpose"`
}
