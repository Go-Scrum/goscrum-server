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

func (service *ProjectService) UpdateAnswerStatus(answer models.Answer) error {
	return service.db.Save(&answer).Error
}

func (service *ProjectService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
	var questions []models.Question

	err := service.db.Find(&questions).Error
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

	for _, question := range questions {
		if !isAnswered(answers, question) {
			return &question, nil
		}
	}
	return nil, nil
}

func isAnswered(answers []models.Answer, question models.Question) bool {
	for _, answer := range answers {
		if answer.QuestionID == question.ID {
			return true
		}
	}
	return false
}