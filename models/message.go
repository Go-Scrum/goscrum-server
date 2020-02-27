package models

type MessageType string

const (
	QuestionMessage   MessageType = "QuestionMessage"
	AnswerMessage     MessageType = "AnswerMessage"
	StandupMessage    MessageType = "StandupMessage"
	GreetingMessage   MessageType = "GreetingMessage"
	OnBoardingMessage MessageType = "OnBoardingMessage"
)

type Message struct {
	Model
	Content       string      `json:"content"`
	UserId        string      `json:"user_id"`
	MessageType   MessageType `json:"message_type"`
	ParticipantID string      `json:"participant_id"`
	Question      Question
	Participant   Participant
}
