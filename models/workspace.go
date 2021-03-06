package models

import "time"

type WorkspaceType string

const (
	Mattermost WorkspaceType = "Mattermost"
)

// TODO end cases
//  - if user has changed the workspace client_id and client_secret details, need to verify login again
// Workspace is used for updating and storing different bot configuration parameters
type Workspace struct {
	Model
	BotUserID     string        `db:"bot_user_id" json:"bot_user_id"`
	Language      string        `db:"language" json:"language"`
	WorkspaceName string        `db:"workspace_name" json:"workspace_name" `
	URL           string        `json:"url"`
	WorkspaceType WorkspaceType `json:"workspace_type"`
	ClientID      string        `json:"client_id"`
	ClientSecret  string        `json:"client_secret"`
	AccessToken   string        `json:"access_token"`
	RefreshToken  string        `json:"refresh_token"`
	Expiry        *time.Time    `json:"expiry"`
	PersonalToken string        `json:"personal_token"`
	Projects      []*Project    `json:"projects"`
	UserEmail     string        `db:"user_email"  json:"user_email"`
}
