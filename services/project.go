package services

import (
	"goscrum/server/models"

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

func (service *ProjectService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
	var questions []models.Question

	err := service.db.Find(&questions).Where("project_id = ?", projectId).Error
	if err != nil {
		return nil, err
	}

	var questionIDs []string
	for _, question := range questions {
		questionIDs = append(questionIDs, question.ID)
	}

	// TODO -- need to add date in this
	var answers []models.Answer
	err = service.db.Find(&answers).Where("question_id in (?) AND participant_id = ?",
		questionIDs,
		participantId,
	).Error
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		for _, answer := range answers {
			if answer.QuestionID != question.ID {
				return &question, nil
			}
		}
	}

	// For question, get the list of integration
	// TODO think for edge cases
	return &questions[0], nil
}
