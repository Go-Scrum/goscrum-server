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
	ParticipantID string `db:"participant_id" json:"participant_id"`
	QuestionID    string `db:"question_id" json:"question_id"`
	Comment       string `db:"comment" json:"comment"`
	BotPostId     string `db:"bot_post_id" json:"bot_post_id"`
	Question      Question
	Participant   Participant
}
