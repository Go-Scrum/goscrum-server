package models

// Project model used for serialization/deserialization stored Projects
type Project struct {
	Model
	WorkspaceID      string         `db:"workspace_id" json:"workspace_id"`
	ChannelName      string         `db:"channel_name" json:"channel_name"`
	ChannelID        string         `db:"channel_id" json:"channel_id"`
	Deadline         string         `db:"deadline" json:"deadline"`
	TZ               string         `db:"tz" json:"tz"`
	OnbordingMessage string         `db:"onbording_message" json:"onbording_message,omitempty"`
	SubmissionDays   string         `db:"submission_days" json:"submission_days,omitempty"`
	Participants     []*Participant `gorm:"many2many:project_participants;"`
	Questions        []*Question    `gorm:"many2many:project_questions;"`
}
