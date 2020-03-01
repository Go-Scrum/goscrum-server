package models

type UserActivityType string

const (
	UserQuestionActivity   UserActivityType = "UserQuestionActivity"
	UserAnswerActivity     UserActivityType = "UserAnswerActivity"
	UserReportActivity     UserActivityType = "UserReportActivity"
	UserStandupActivity    UserActivityType = "UserStandupActivity"
	UserGreetingActivity   UserActivityType = "UserGreetingActivity"
	UserOnBoardingActivity UserActivityType = "UserOnBoardingActivity"
)

type UserActivity struct {
	Model
	UserId        string           `json:"user_id"`
	ChannelID     string           `json:"channel_id"`
	ProjectID     string           `db:"project_id" json:"project_id"`
	ParticipantID string           `db:"participant_id" json:"participant_id"`
	QuestionID    string           `db:"question_id" json:"question_id"`
	BotPostId     string           `db:"bot_post_id" json:"bot_post_id"`
	AnswerID      string           `db:"answer_id" json:"answer_id"`
	ActivityType  UserActivityType `json:"activity_type"`
	Question      Question
	Participant   Participant
	Project       Project
}
