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
	// TODO get list of participants by user_Id
	//  merge with projects

	var participants []models.Participant
	var userIds []string
	for _, participant := range project.Participants {
		userIds = append(userIds, participant.UserID)
	}
	service.db.Where("user_id in (?)", userIds).Find(&participants)
	for i, user := range project.Participants {
		for _, participant := range participants {
			if user.UserID == participant.UserID {
				project.Participants[i] = &participant
				break
			}
		}
	}
	err := service.db.Save(&project).Error
	return project, err
}

func (service *ProjectService) GetAll(workspaceId string) ([]models.Project, error) {
	var projects []models.Project

	if err := service.db.Where("workspace_id = ?", workspaceId).Find(&projects).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return projects, err
	}

	return projects, nil
}

func (service *ProjectService) GetByID(id string) (models.Project, error) {
	var project models.Project

	err := service.db.
		Preload("Questions").
		Preload("Participants").
		Where("id = ?", id).First(&project).Error

	return project, err
}

func (service *ProjectService) GetProjectById(id string) (models.Project, error) {
	var project models.Project

	err := service.db.
		Where("id = ?", id).First(&project).Error

	return project, err
}

func (service *ProjectService) UpdateAnswerPostId(answer models.Answer) error {
	existingAnswer := models.Answer{}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	err := service.db.Where("question_id = (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
		answer.QuestionID,
		answer.ParticipantID,
		yesterday,
		today,
	).First(&existingAnswer).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return service.db.Save(&answer).Error
	}
	return service.db.Save(&existingAnswer).Error
}

func (service *ProjectService) AddUserAnswer(answer models.Answer) (models.Answer, error) {
	err := service.db.Create(&answer).Error
	return answer, err
}

func (service *ProjectService) GetProjectQuestions(projectId string) ([]models.Question, error) {
	var questions []models.Question
	err := service.db.
		Where("project_id = ?", projectId).
		Order("sequence asc").
		Find(&questions).Error
	return questions, err
}

func (service *ProjectService) GetParticipantsAnswer(answerIds []string) ([]models.Answer, error) {
	var answers []models.Answer
	err := service.db.
		Preload("Question").
		Order("created_at asc").
		Where("id in (?)", answerIds).
		Find(&answers).Error
	return answers, err
}

//
//func (service *ProjectService) UserInteraction(userId string, message models.Message, workspace models.Workspace) (*models.Message, error) {
//	participant := models.Participant{}
//	err := service.db.
//		Where("user_id = ?", userId).
//		First(&participant).Error
//
//	if err != nil && !gorm.IsRecordNotFoundError(err) {
//		// TODO- send message that user is not configured for goscrum.io
//		return nil, err
//	}
//	activity, err := UserActivityService{}
//	// TODO -- get user last activities group
//
//	// TODO if last activity type = question then consider this as answer
//
//	// TODO if last activity type = report then ask for restart question
//
//	// TODO - need to check for workspace project name too here for extra validation
//	participant := models.Participant{}
//	err := service.db.
//		Where("user_id = ?", userId).
//		First(&participant).Error
//
//	if err != nil && !gorm.IsRecordNotFoundError(err) {
//		return nil, err
//	}
//
//	if gorm.IsRecordNotFoundError(err) {
//		// TODO -- send message to asked for admin to configure for your channels
//		return nil, err
//	}
//
//	if strings.ToLower(message.Content) == "cancel" || strings.ToLower(message.Content) == "cancelled" {
//		// TODO -cancel standup
//	}
//
//	existingAnswer := models.Answer{}
//	today := time.Now()
//	yesterday := today.AddDate(0, 0, -1)
//	err = service.db.
//		Preload("Question").
//		Where("participant_id = ? AND updated_at BETWEEN ? AND ?",
//			participant.ID,
//			yesterday,
//			today,
//		).
//		Order("updated_at desc").
//		First(&existingAnswer).Error
//
//	if err != nil && !gorm.IsRecordNotFoundError(err) {
//		return nil, err
//	}
//	if gorm.IsRecordNotFoundError(err) {
//		// TODO need to check when record not found
//		return nil, nil
//	}
//	existingAnswer.Comment = message.Content
//	existingAnswer.ParticipantID = participant.ID
//	err = service.db.Save(&existingAnswer).Error
//
//	if err != nil {
//		return nil, err
//	}
//
//	var questions []models.Question
//	err = service.db.Where("project_id = ?", existingAnswer.Question.ProjectId).Find(&questions).Error
//	if err != nil {
//		return nil, err
//	}
//
//	var questionIDs []string
//	for _, question := range questions {
//		questionIDs = append(questionIDs, question.ID)
//	}
//
//	today = time.Now()
//	yesterday = today.AddDate(0, 0, -1)
//	var answers []models.Answer
//	err = service.db.Where("question_id in (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
//		questionIDs,
//		participant.ID,
//		yesterday,
//		today,
//	).Find(&answers).Error
//	if err != nil {
//		return nil, err
//	}
//
//	question := models.Question{}
//	for _, ques := range questions {
//		if !isAnswered(answers, ques) {
//			question = ques
//			break
//		}
//	}
//	if question.Title == "" {
//		var project models.Project
//		err = service.db.Where("id = ?", existingAnswer.Question.ProjectId).First(&project).Error
//		if err != nil {
//			return nil, err
//		}
//		var attachments []*model.SlackAttachment
//		for _, answer := range answers {
//			attachments = append(attachments, &model.SlackAttachment{
//				Color:   "#FF0000",
//				Pretext: answer.Comment,
//				Title:   answer.Question.Title,
//			})
//		}
//
//		return &models.Message{
//			Attachments:   attachments,
//			Content:       question.Title,
//			UserId:        userId,
//			MessageType:   models.StandupMessage,
//			ParticipantID: participant.ID,
//			ChannelId:     project.ChannelID,
//		}, nil
//	}
//
//	return &models.Message{
//		Content:       question.Title,
//		UserId:        userId,
//		MessageType:   models.QuestionMessage,
//		ParticipantID: participant.ID,
//		Question:      question,
//	}, nil
//}

//func (service *ProjectService) GetParticipantQuestion(projectId, participantId string) (*models.Question, error) {
//	var questions []models.Question
//
//	err := service.db.Where("project_id = ?", projectId).Find(&questions).Error
//	if err != nil {
//		return nil, err
//	}
//
//	var questionIDs []string
//	for _, question := range questions {
//		questionIDs = append(questionIDs, question.ID)
//	}
//
//	var answers []models.Answer
//	today := time.Now()
//	yesterday := today.AddDate(0, 0, -1)
//	err = service.db.Where("question_id in (?) AND participant_id = ? AND updated_at BETWEEN ? AND ?",
//		questionIDs,
//		participantId,
//		yesterday,
//		today,
//	).Find(&answers).Error
//	if err != nil {
//		return nil, err
//	}
//
//	question := models.Question{}
//	for _, ques := range questions {
//		if !isAnswered(answers, ques) {
//			question = ques
//			break
//		}
//	}
//	return &question, nil
//}

func (service *ProjectService) GetQuestionDetails(questionId string) (*models.Question, error) {
	question := models.Question{}

	err := service.db.Where("id = ?", questionId).First(&question).Error
	if err != nil {
		return nil, err
	}

	// Need to call integration to get more information about the question
	return &question, nil
}

func (service *ProjectService) UpdateParticipant(participant models.Participant) error {
	return service.db.Save(&participant).Error
}

//func isAnswered(answers []models.Answer, question models.Question) bool {
//	for _, answer := range answers {
//		if answer.QuestionID == question.ID {
//			return true
//		}
//	}
//	return false
//}
