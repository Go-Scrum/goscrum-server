package services

type BotService struct {
	workspaceService  WorkspaceService
	projectService    ProjectService
	mattermostService MattermostService
}

func NewBotService(workspaceService WorkspaceService, projectService ProjectService, mattermostService MattermostService) BotService {
	return BotService{
		workspaceService:  workspaceService,
		projectService:    projectService,
		mattermostService: mattermostService,
	}
}
