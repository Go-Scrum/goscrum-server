package services

import (
	"goscrum/server/models"

	"github.com/jinzhu/gorm"
)

type WorkspaceService struct {
	db *gorm.DB
}

func NewWorkspaceService(db *gorm.DB) WorkspaceService {
	return WorkspaceService{db: db}
}

//CreateWorkspace creates bot properties for the newly created bot
func (service *WorkspaceService) Save(workspace models.Workspace) (models.Workspace, error) {
	//err := bs.Validate()
	//if err != nil {
	//	return bs, err
	//}
	err := service.db.Save(&workspace).Error
	return workspace, err
}

//GetAllWorkspaces returns all workspaces stored in DB
func (service *WorkspaceService) GetAllWorkspaces() ([]models.Workspace, error) {
	var workspaces []models.Workspace
	err := service.db.
		Preload("Projects.Participants").
		Preload("Projects.Questions").
		Find(&workspaces).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return workspaces, err
	}

	return workspaces, nil
}

//GetWorkspaceByWorkspaceID returns a particular bot
func (service *WorkspaceService) GetWorkspaceByWorkspaceID(workspaceID string) (models.Workspace, error) {
	workspace := models.Workspace{}
	err := service.db.
		Where("workspace_id = ?", workspaceID).
		First(&workspace).Error

	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

//GetWorkspaceByBotAccessToken returns a particular bot
func (service *WorkspaceService) GetWorkspaceByBotAccessToken(botAccessToken string) (models.Workspace, error) {
	workspace := models.Workspace{}
	err := service.db.
		Where("bot_access_token = ?", botAccessToken).
		First(&workspace).Error

	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

func (service *WorkspaceService) GetWorkspaceByUserEmail(email string) (*models.Workspace, error) {
	var workspace models.Workspace
	err := service.db.
		Where("user_email = ?", email).
		First(&workspace).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return &workspace, err
	}

	return &workspace, nil
}

//GetWorkspace returns a particular bot
func (service *WorkspaceService) GetWorkspace(id string) (models.Workspace, error) {
	workspace := models.Workspace{}
	err := service.db.
		Where("id = ?", id).
		First(&workspace).Error

	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

//GetWorkspace returns a particular bot
func (service *WorkspaceService) GetWorkspaceByToken(token string) (models.Workspace, error) {
	workspace := models.Workspace{}
	err := service.db.
		Preload("Projects.Participants").
		Preload("Projects.Questions", func(db *gorm.DB) *gorm.DB {
			return db.Order("questions.sequence ASC")
		}).
		Where("personal_token = ?", token).
		First(&workspace).Error

	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

//DeleteWorkspaceByID deletes bot
func (service *WorkspaceService) DeleteWorkspaceByID(id string) error {
	existingWorkspace := &models.Workspace{}
	err := service.db.
		Where("id = ?", id).
		First(existingWorkspace).Error

	if err != nil {
		return err
	}

	return service.db.Delete(&existingWorkspace).Error
}

//DeleteWorkspace deletes bot
func (service *WorkspaceService) DeleteWorkspace(workspaceID string) error {
	existingWorkspace := &models.Workspace{}
	err := service.db.
		Where("workspace_id = ?", workspaceID).
		First(existingWorkspace).Error

	if err != nil {
		return err
	}

	return service.db.Delete(&existingWorkspace).Error
}
