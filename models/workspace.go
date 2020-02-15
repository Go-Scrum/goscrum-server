package models

type WorkspaceType int

const (
	Mattermost WorkspaceType = 0
)

// Workspace is used for updating and storing different bot configuration parameters
type Workspace struct {
	Model
	BotUserID              string        `db:"bot_user_id" json:"bot_user_id"`
	NotifierInterval       int           `db:"notifier_interval" json:"notifier_interval" `
	Language               string        `db:"language" json:"language" `
	MaxReminders           int           `db:"max_reminders" json:"max_reminders" `
	ReminderOffset         int           `db:"reminder_offset" json:"reminder_offset" `
	BotAccessToken         string        `db:"bot_access_token" json:"bot_access_token" `
	AppId                  string        `db:"app_id" json:"app_id"`
	WorkspaceName          string        `db:"workspace_name" json:"workspace_name" `
	ReportingChannel       string        `db:"reporting_channel" json:"reporting_channel"`
	ReportingTime          string        `db:"reporting_time" json:"reporting_time"`
	ProjectsReportsEnabled bool          `db:"projects_reports_enabled" json:"projects_reports_enabled"`
	URL                    string        `json:"url"`
	AuthType               WorkspaceType `json:"auth_type"`
	ClientID               string        `json:"client_id"`
	ClientSecret           string        `json:"client_secret"`
	AccessToken            string        `json:"access_token"`
	Projects               []*Project
}
