package services

type BotService struct {
	workspaceService WorkspaceService
}

func NewBotService(workspaceService WorkspaceService) BotService {
	return BotService{
		workspaceService: workspaceService,
	}
}
