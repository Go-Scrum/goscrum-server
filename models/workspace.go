package models

import "time"

type WorkspaceType int

const (
	Mattermost WorkspaceType = 0
)

// TODO end cases
//  - if user has changed the workspace client_id and client_secret details, need to verify login again
// Workspace is used for updating and storing different bot configuration parameters
type Workspace struct {
	Model
	Language      string        `db:"language" json:"language"`
	WorkspaceName string        `db:"workspace_name" json:"workspace_name" `
	URL           string        `json:"url"`
	WorkspaceType WorkspaceType `json:"workspace_type"`
	ClientID      string        `json:"client_id"`
	ClientSecret  string        `json:"client_secret"`
	AccessToken   string        `json:"access_token"`
	RefreshToken  string        `json:"refresh_token"`
	Expiry        *time.Time    `json:"expiry,omitempty"`
	Projects      []*Project
}

