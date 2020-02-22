package models

type QuestionType int

const (
	Text       QuestionType = 0
	Numeric    QuestionType = 1
	PreDefined QuestionType = 2
)

type Question struct {
	Model
	Title     string
	Type      QuestionType
	ProjectId string
}
