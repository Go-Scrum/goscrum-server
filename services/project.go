package services

import (
	"goscrum/server/models"
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
	existingAnswer.Comment = answer.Comment
	return service.db.Save(&existingAnswer).Error
}

func (service *ProjectService) UserMessage(userId string, answer models.Answer) (*models.Answer, error) {
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

	existingAnswer := models.Answer{}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	err = service.db.
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
	existingAnswer.Comment = answer.Comment
	existingAnswer.ParticipantID = participant.ID
	err = service.db.Save(&existingAnswer).Error
	return &existingAnswer, nil
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
