package main

import (
	"context"
	"fmt"
	"goscrum/server/db"
	"goscrum/server/models"
	"goscrum/server/services"
	"time"

	"github.com/1set/cronrange"
	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	botUsername    = "GoScrum"
	botDisplayName = "GoScrum"
	botDescription = "A bot account created by the GoScrum plugin."
)

func HandleRequest(_ context.Context) {
	fmt.Println("Bot is running")
	db := db.DbClient(true)

	defer db.Close()

	workspaceService := services.NewWorkspaceService(db)
	projectService := services.NewProjectService(db)
	//mattermostService := services.NewMattermostService(workspaceService, projectService)
	//botService := services.NewBotService(workspaceService, projectService, mattermostService)

	workspaces, err := workspaceService.GetAllWorkspaces()
	if err != nil {
		fmt.Println(err.Error())
		// TODO -- write error message
	}
	now := time.Now()

	for _, workspace := range workspaces {
		// TODO -- only for mattermost
		//if workspace.WorkspaceType == models.Mattermost {
		// TODO -- put more validation if required
		if workspace.AccessToken != "" && workspace.Projects != nil || len(workspace.Projects) > 0 {
			// TODO -- create bot at workspace level VIMP
			apiClient := model.NewAPIv4Client(workspace.URL)
			apiClient.SetOAuthToken(workspace.AccessToken)
			//bot, res := apiClient.GetBot(workspace.BotUserID, "")
			//if res != nil {
			//	fmt.Println(res.Error.Message))
			//	// TODO -- write error message
			//}
			//if bot != nil {
			//	systemBot := &model.Bot{
			//		Username:    botUsername,
			//		DisplayName: botDisplayName,
			//		Description: botDescription,
			//	}
			//	bot, err = apiClient.CreateBot(systemBot)
			//	if err != nil {
			//		// TODO -- write error message
			//		continue
			//	}
			//
			//	workspace.BotUserID = bot.UserId
			//	_, err := workspaceService.Save(workspace)
			//	if err != nil {
			//		// TODO -- write error message
			//		continue
			//	}
			//}

			for _, project := range workspace.Projects {
				if cr, err := cronrange.ParseString(project.ReportingTime); err == nil {
					if cr.IsWithin(now) {
						for _, participant := range project.Participants {
							message := fmt.Sprintf("Hello %s :wave: It's time for **%s** in # %s\n Please share what you've been working on",
								participant.RealName,
								project.Name,
								project.ChannelName,
							)
							post := model.Post{
								// TODO - format the name based on first name and lastName
								Message: message,
								UserId:  workspace.BotUserID,
							}
							if participant.BotChannelID != "" {
								channel, res := apiClient.CreateDirectChannel(workspace.BotUserID, participant.UserID)
								if res.StatusCode != 201 {
									fmt.Println(res.Error.Message)
									// TODO log error message
									break
								}
								participant.BotChannelID = channel.Id
								err := projectService.UpdateParticipant(*participant)
								if err != nil {
									fmt.Println(err.Error())
									// TODO -- write error message
									continue
								}
							}
							post.ChannelId = participant.BotChannelID
							_, res := apiClient.CreatePost(&post)
							if res.StatusCode != 201 {
								fmt.Println(res.Error.Message)
								// TODO log error message
								continue
							}
							question, err := projectService.GetQuestionDetails(project.Questions[0].ID)
							if err != nil {
								fmt.Println(err.Error())
								// TODO -- write error message
								continue
							}

							fmt.Println("Question", question.ID)
							if question.Title != "" {
								post = model.Post{
									Message:   question.Title,
									UserId:    workspace.BotUserID,
									ChannelId: participant.BotChannelID,
								}

								createdPost, res := apiClient.CreatePost(&post)
								if res.StatusCode != 201 {
									// TODO log error message
									continue
								}
								projectService.UpdateAnswerPostId(models.Answer{
									ParticipantID: participant.ID,
									QuestionID:    question.ID,
									Comment:       "",
									BotPostId:     createdPost.Id,
								})
							}
						}
					}
				}

			}
		}
		//}
	}
}

func main() {
	//lambda.Start(HandleRequest)
	HandleRequest(context.Background())
}
