package services

import (
	"goscrum/server/models"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) ProjectService {
	return ProjectService{db: db}
}

func (service *ProjectService) Save(project models.Project) (models.Project, error) {
	err := service.db.Save(&project).Error
	return project, err
}

func (service *ProjectService) GetAll() ([]models.Project, error) {
	var projects []models.Project

	if err := service.db.Find(&projects).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return projects, err
	}

	return projects, nil
}

func (service *ProjectService) GetByID(id string) (models.Project, error) {
	var project models.Project

	err := service.db.Where("id = ?", id).First(&project).Error

	return project, err
}

func (service *ProjectService) UpdateAnswerPostId(answer models.Answer) error {
	existingAnswer := models.Answer{}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	err := service.db.Where("question_id = (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
		answer.QuestionID,
		answer.ParticipantID,
		yesterday,
		today,
	).First(&existingAnswer).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return service.db.Save(&answer).Error
	}
	return service.db.Save(&existingAnswer).Error
}

func (service *ProjectService) UserInteraction(userId string, message models.Message) (*models.Message, error) {
	// TODO - need to check for workspace project name too here for extra validation
	participant := models.Participant{}
	err := service.db.
		Where("user_id = ?", userId).
		First(&participant).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		// TODO -- send message to asked for admin to configure for your channels
		return nil, err
	}

	if strings.ToLower(message.Content) == "cancel" || strings.ToLower(message.Content) == "cancelled" {
		// TODO -cancel standup
	}

	existingAnswer := models.Answer{}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	err = service.db.
		Preload("Question").
		Where("participant_id = ? AND updated_at BETWEEN ? AND ?",
			participant.ID,
			yesterday,
			today,
		).
		Order("updated_at desc").
		First(&existingAnswer).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		// TODO need to check when record not found
		return nil, nil
	}
	existingAnswer.Comment = message.Content
	existingAnswer.ParticipantID = participant.ID
	err = service.db.Save(&existingAnswer).Error

	if err != nil {
		return nil, err
	}

	var questions []models.Question
	err = service.db.Where("project_id = ?", existingAnswer.Question.ProjectId).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	var questionIDs []string
	for _, question := range questions {
		questionIDs = append(questionIDs, question.ID)
	}

	today = time.Now()
	yesterday = today.AddDate(0, 0, -1)
	var answers []models.Answer
	err = service.db.Where("question_id in (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
		questionIDs,
		participant.ID,
		yesterday,
		today,
	).Find(&answers).Error
	if err != nil {
		return nil, err
	}

	question := models.Question{}
	for _, ques := range questions {
		if !isAnswered(answers, ques) {
			question = ques
			break
		}
	}
	if question.Title == "" {
		// TODO -- all questions are done - send completion message
		return nil, nil
	}
	return &models.Message{
		Content:       question.Title,
		UserId:        userId,
		MessageType:   models.QuestionMessage,
		ParticipantID: participant.ID,
		Question:      question,
	}, nil
}

func (service *ProjectService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
	var questions []models.Question

	err := service.db.Where("project_id = ?", projectId).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	var questionIDs []string
	for _, question := range questions {
		questionIDs = append(questionIDs, question.ID)
	}

	var answers []models.Answer
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	err = service.db.Where("question_id in (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
		questionIDs,
		participantId,
		yesterday,
		today,
	).Find(&answers).Error
	if err != nil {
		return nil, err
	}

	question := models.Question{}
	for _, ques := range questions {
		if !isAnswered(answers, ques) {
			question = ques
			break
		}
	}
	return &question, nil
}

func (service *ProjectService) GetQuestionDetails(questionId string) (*models.Question, error) {
	question := models.Question{}

	err := service.db.Where("id = ?", questionId).First(&question).Error
	if err != nil {
		return nil, err
	}

	// Need to call integration to get more information about the question
	return &question, nil
}

func isAnswered(answers []models.Answer, question models.Question) bool {
	for _, answer := range answers {
		if answer.QuestionID == question.ID {
			return true
		}
	}
	return false
}
