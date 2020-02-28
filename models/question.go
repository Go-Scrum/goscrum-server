package models

type QuestionType string

const (
	Text       QuestionType = "Text"
	Numeric    QuestionType = "Numeric"
	PreDefined QuestionType = "PreDefined"
)

type Question struct {
	Model
	Title     string
	Type      QuestionType
	ProjectId string
}
