package models

import "github.com/mattermost/mattermost-server/v5/model"

type MessageType string

const (
	QuestionMessage   MessageType = "QuestionMessage"
	ErrorMessage      MessageType = "ErrorMessage"
	AnswerMessage     MessageType = "AnswerMessage"
	StandupMessage    MessageType = "StandupMessage"
	GreetingMessage   MessageType = "GreetingMessage"
	OnBoardingMessage MessageType = "OnBoardingMessage"
)

type Message struct {
	Model
	Attachments   []*model.SlackAttachment
	Content       string      `json:"content"`
	UserId        string      `json:"user_id"`
	ChannelId     string      `json:"channel_id"`
	MessageType   MessageType `json:"message_type"`
	ParticipantID string      `json:"participant_id"`
	Question      Question
	Participant   Participant
}
