package services

import (
	"goscrum/server/models"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mattermost/mattermost-server/v5/model"
)

type MattermostService struct {
	workspaceService WorkspaceService
}

func NewMattermostService(workspaceService WorkspaceService) MattermostService {
	return MattermostService{workspaceService: workspaceService}
}

func (m *MattermostService) GetAllPublicChannels(workspaceId string) ([]models.Channel, error) {
	workspace, err := m.workspaceService.GetWorkspace(workspaceId)
	var channels []models.Channel

	if err != nil {
		return channels, err
	}
	apiClient := model.NewAPIv4Client(workspace.URL)
	apiClient.SetOAuthToken(workspace.AccessToken)

	// TODO -- need to work on pagination
	mattermostChannels, _, res := apiClient.GetAllChannelsWithCount(0, 100, "")

	if res.StatusCode == 200 {
		for _, ch := range *mattermostChannels {
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
	return nil, http.ErrNotSupported
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

func (m *MattermostService) Standup(personalToken, channelId, userId, botId string) error {
	workspace, err := m.workspaceService.GetWorkspaceByToken(personalToken)
	if err != nil {
		return err
	}
	apiClient := model.NewAPIv4Client(workspace.URL)
	apiClient.SetOAuthToken(workspace.AccessToken)

	post, res := apiClient.CreatePost(&model.Post{
		Id:            "",
		CreateAt:      time.Now().Unix(),
		IsPinned:      false,
		UserId:        botId,
		ChannelId:     channelId,
		Message:       "Hello Duragaprasad",
		MessageSource: "",
	})

	spew.Dump(post)

	if res.StatusCode != 200 {
		return res.Error
	}
	return nil
}
