package services

import (
	"goscrum/server/models"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

type MattermostService struct {
	workspaceService WorkspaceService
	projectService   ProjectService
}

func NewMattermostService(workspaceService WorkspaceService, projectService ProjectService) MattermostService {
	return MattermostService{workspaceService: workspaceService, projectService: projectService}
}

func (m *MattermostService) GetPublicChannelsForTeam(workspaceId, teamId string) ([]models.Channel, error) {
	workspace, err := m.workspaceService.GetWorkspace(workspaceId)
	var channels []models.Channel

	if err != nil {
		return channels, err
	}
	apiClient := model.NewAPIv4Client(workspace.URL)
	apiClient.SetOAuthToken(workspace.AccessToken)

	// TODO -- need to work on pagination
	mattermostChannels, res := apiClient.GetPublicChannelsForTeam(teamId, 0, 100, "")

	if res.StatusCode == 200 {
		for _, ch := range mattermostChannels {
			channels = append(channels, models.Channel{
				Id:          ch.Id,
				TeamId:      ch.TeamId,
				Type:        ch.Type,
				DisplayName: ch.DisplayName,
				Name:        ch.Name,
				Purpose:     ch.Purpose,
			})
		}
		return channels, nil
	}
	// TODO edge case, when token is expired
	return nil, res.Error
}

func (m *MattermostService) GetParticipants(workspaceId, channelId string) ([]models.Participant, error) {
	workspace, err := m.workspaceService.GetWorkspace(workspaceId)
	var participants []models.Participant

	if err != nil {
		return participants, err
	}
	apiClient := model.NewAPIv4Client(workspace.URL)
	apiClient.SetOAuthToken(workspace.AccessToken)

	// TODO -- need to work on pagination
	users, res := apiClient.GetUsersInChannel(channelId, 0, 100, "")

	if res.StatusCode == 200 {
		for _, user := range users {
			participants = append(participants, models.Participant{
				Email:       user.Email,
				UserID:      user.Id,
				WorkspaceID: workspaceId,
				Role:        user.Roles,
				RealName:    user.Username,
				FirstName:   user.FirstName,
				LastName:    user.LastName,
			})
		}
		return participants, nil
	}
	// TODO edge case, when token is expired
	return nil, res.Error
}

func (m *MattermostService) GetAllTeams(workspaceId string) ([]models.Team, error) {
	workspace, err := m.workspaceService.GetWorkspace(workspaceId)
	var teams []models.Team

	if err != nil {
		return teams, err
	}
	apiClient := model.NewAPIv4Client(workspace.URL)
	apiClient.SetOAuthToken(workspace.AccessToken)

	// TODO -- need to work on pagination
	mattermostTeams, _, res := apiClient.GetAllTeamsWithTotalCount("", 0, 100)

	if res.StatusCode == 200 {
		for _, mTeam := range mattermostTeams {
			teams = append(teams, models.Team{
				Id:          mTeam.Id,
				DisplayName: mTeam.DisplayName,
				Name:        mTeam.Name,
				Description: mTeam.Description,
			})
		}
		return teams, nil
	}
	// TODO edge case, when token is expired
	return nil, http.ErrNotSupported
}

//GetWorkspace returns a particular bot
func (m *MattermostService) GetWorkspaceByToken(token string) (models.Workspace, error) {
	return m.workspaceService.GetWorkspaceByToken(token)
}

//GetWorkspace returns a particular bot
func (m *MattermostService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
	return m.projectService.GetParticipantQuestion(projectId, participantId)
}

func (m *MattermostService) UpdateAnswerPostId(answer models.Answer) error {
	return m.projectService.UpdateAnswerPostId(answer)
}

func (m *MattermostService) UserInteraction(userId string, message models.Message) (*models.Message, error) {
	return m.projectService.UserInteraction(userId, message)
}

func (m *MattermostService) GetQuestionDetails(questionId string) (*models.Question, error) {
	return m.projectService.GetQuestionDetails(questionId)
}
