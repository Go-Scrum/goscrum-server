package models

type AnswerStatus string

const (
	Asked     AnswerStatus = "asked"
	Answered  AnswerStatus = "answered"
	Cancelled AnswerStatus = "cancelled"
)

// Standup model used for serialization/deserialization stored standups
type Answer struct {
	Model
	WorkspaceID string       `db:"workspace_id" json:"workspace_id"`
	ChannelID   string       `db:"channel_id" json:"channel_id"`
	UserID      string       `db:"user_id" json:"user_id"`
	QuestionID  string       `db:"question_id" json:"question_id"`
	Comment     string       `db:"comment" json:"comment"`
	Status      AnswerStatus `db:"status" json:"status"`
	MessageTS   string       `db:"message_ts" json:"message_ts"`
}
