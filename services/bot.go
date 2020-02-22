package services

import "goscrum/server/models"

type BotService struct {
	workspaceService WorkspaceService
}

func NewBotService(workspaceService WorkspaceService) BotService {
	return BotService{workspaceService: workspaceService}
}

func (b *BotService) GetAllProjects(workspaceId string) ([]models.Project, error) {
	return nil, nil
}
