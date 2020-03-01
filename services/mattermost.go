package services

import (
	"fmt"
	"goscrum/server/constants"
	"goscrum/server/models"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

type MattermostService struct {
	workspaceService    WorkspaceService
	projectService      ProjectService
	userActivityService UserActivityService
	participantService  ParticipantService
}

func NewMattermostService(
	workspaceService WorkspaceService,
	projectService ProjectService,
	userActivityService UserActivityService,
	participantService ParticipantService,
) MattermostService {
	return MattermostService{
		workspaceService:    workspaceService,
		projectService:      projectService,
		userActivityService: userActivityService,
		participantService:  participantService,
	}
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
//func (m *MattermostService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
//	return m.projectService.GetParticipantQuestion(projectId, participantId)
//}

func (m *MattermostService) AddUserActivity(userActivity models.UserActivity) error {
	return m.userActivityService.Add(userActivity)
}

func (m *MattermostService) UserInteraction(userId string, message models.Message) (*models.Message, error) {
	participant, err := m.participantService.GetParticipantByUserId(userId)
	if err != nil {
		return &models.Message{
			Content:       err.Error(),
			UserId:        userId,
			MessageType:   models.QuestionMessage,
			ParticipantID: participant.ID,
		}, nil
	}
	if participant == nil {
		// TODO send message participant is not configured
	}

	userActivity, err := m.userActivityService.GetLastUserActivity(participant.ID)
	if err != nil {
		return &models.Message{
			Content:       err.Error(),
			UserId:        userId,
			MessageType:   models.ErrorMessage,
			ParticipantID: participant.ID,
		}, nil
	}

	if userActivity == nil {
		var options []*model.PostActionOptions
		for _, project := range participant.Projects {
			options = append(options, &model.PostActionOptions{
				Text:  project.Name,
				Value: project.ID,
			})
		}
		options = append(options, &model.PostActionOptions{
			Text:  "Skip",
			Value: "Skip",
		})
		message := fmt.Sprintf("Hello @%s :wave: Would you like to start GoScrum.io?",
			participant.RealName,
		)
		attachment := model.SlackAttachment{
			Color: "#00FFFF", // TODO later change the color
			Text:  message,

			Actions: []*model.PostAction{
				&model.PostAction{
					Type:       model.POST_ACTION_TYPE_BUTTON,
					Name:       "Select an option...",
					Disabled:   false,
					DataSource: "",
					//Options:       options,
					DefaultOption: "Skip",
					Integration: &model.PostActionIntegration{
						URL: fmt.Sprintf("/plugins/%s/api/v1/user/action", constants.MattermostPluginId),
						Context: map[string]interface{}{
							"action": "action",
						},
					},
				},
			},
		}
		return &models.Message{
			Attachments:   []*model.SlackAttachment{&attachment},
			UserId:        participant.UserID,
			MessageType:   models.StandupMessage,
			ParticipantID: "",
		}, nil
	}

	if userActivity.ActivityType == models.UserQuestionActivity {
		savedAnswer, err := m.projectService.AddUserAnswer(models.Answer{
			ParticipantID: participant.ID,
			QuestionID:    userActivity.QuestionID,
			Comment:       message.Content,
			BotPostId:     "",
		})

		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}

		err = m.userActivityService.Add(models.UserActivity{
			UserId:        userId,
			ChannelID:     userActivity.ChannelID,
			ProjectID:     userActivity.ProjectID,
			ParticipantID: participant.ID,
			QuestionID:    userActivity.QuestionID,
			AnswerID:      savedAnswer.ID,
			BotPostId:     "",
			ActivityType:  models.UserAnswerActivity,
		})

		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}

		// get list of questions for project
		questions, err := m.projectService.GetProjectQuestions(userActivity.ProjectID)
		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}

		userActivities, err := m.userActivityService.GetUserActivities(participant.ID)
		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}
		var newQuestion *models.Question
		for _, question := range questions {
			if !isAnswered(userActivities, question) {
				newQuestion = &question
				break
			}
		}
		if newQuestion != nil {
			// send next question
			return &models.Message{
				Content:       newQuestion.Title,
				UserId:        userId,
				MessageType:   models.QuestionMessage,
				ParticipantID: participant.ID,
				Question:      *newQuestion,
			}, nil
		}

		var answerIDs []string
		for _, question := range questions {
			for _, activity := range userActivities {
				if activity.ActivityType == models.UserAnswerActivity &&
					activity.QuestionID == question.ID {
					answerIDs = append(answerIDs, activity.AnswerID)
					break
				}
			}
		}
		projectAnswers, err := m.projectService.GetParticipantsAnswer(answerIDs)
		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}

		var attachments []*model.SlackAttachment
		for _, answer := range projectAnswers {
			attachments = append(attachments, &model.SlackAttachment{
				Color: answer.Question.Color,
				Text:  fmt.Sprintf("%s\r\n%s", answer.Question.Title, answer.Comment),
			})
		}

		project, err := m.projectService.GetProjectById(userActivity.ProjectID)
		if err != nil {
			return &models.Message{
				Content:       err.Error(),
				UserId:        userId,
				MessageType:   models.ErrorMessage,
				ParticipantID: participant.ID,
			}, nil
		}
		return &models.Message{
			Content:       fmt.Sprintf("@%s updates", participant.RealName),
			Attachments:   attachments,
			UserId:        userId,
			MessageType:   models.ReportMessage,
			ParticipantID: participant.ID,
			ChannelId:     project.ChannelID,
		}, nil
		// TODO -- all questions are anwsered
	}

	// TODO -- needto check for this condition
	return nil, nil
	//return m.projectService.UserInteraction(userId, message, workspace)
}

func isAnswered(userActivities []models.UserActivity, question models.Question) bool {
	for _, activity := range userActivities {
		if activity.ActivityType == models.UserAnswerActivity &&
			activity.QuestionID == question.ID {
			return true
		}
	}
	return false
}

func (m *MattermostService) GetQuestionDetails(questionId string) (*models.Question, error) {
	return m.projectService.GetQuestionDetails(questionId)
}
