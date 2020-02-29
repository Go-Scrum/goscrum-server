package services

import (
	"github.com/xanzy/go-gitlab"
	"goscrum/server/constants"
	"goscrum/server/models"
	"os"
	"time"
)

type GitlabClient struct{}

func (GitlabClient) Client() *gitlab.Client {
	return gitlab.NewClient(nil, os.Getenv(constants.GitlabAccessToken))
}

type GitlabService struct {
	client *GitlabClient
}

func NewGitlabPService() *GitlabService {
	return &GitlabService{client: &GitlabClient{}}
}

func (service GitlabService) Commits(req models.RequestCommits, id string) ([]*gitlab.Commit, error) {
	commits, _, err := service.client.Client().Commits.ListCommits(id, &gitlab.ListCommitsOptions{Since: gitlab.Time(req.Since)})
	return commits, err
}

func (service GitlabService) Issues(req models.RequestIssues) ([]*gitlab.Issue, error) {
	issues, _, err := service.client.Client().Issues.ListIssues(&gitlab.ListIssuesOptions{
		ListOptions:  gitlab.ListOptions{},
		State:        gitlab.String(req.State),
		AuthorID:     gitlab.Int(req.AuthorID),
		AssigneeID:   gitlab.Int(req.AssigneeID),
		Search:       gitlab.String(req.Search),
		UpdatedAfter: gitlab.Time(req.UpdatedAfter),
	})
	return issues, err
}

func (service GitlabService) Users(search string) ([]*gitlab.User, error) {
	users, _, err := service.client.Client().Users.ListUsers(&gitlab.ListUsersOptions{Search: gitlab.String(search)})
	return users, err
}

func (service GitlabService) Projects(search string) ([]*gitlab.Project, error) {
	projects, _, err := service.client.Client().Projects.ListProjects(&gitlab.ListProjectsOptions{
		Search:     gitlab.String(search),
		Visibility: gitlab.Visibility(gitlab.PrivateVisibility),
	})
	return projects, err
}

func (service GitlabService) UserContributions(userId, action string, after time.Time) ([]*gitlab.ContributionEvent, error) {
	var isoTime gitlab.ISOTime
	isoTime = gitlab.ISOTime(after)
	var eventAction gitlab.EventTypeValue
	eventAction = gitlab.EventTypeValue(action)
	events, _, err := service.client.Client().Users.ListUserContributionEvents(userId, &gitlab.ListContributionEventsOptions{
		Action: &eventAction,
		After:  &isoTime,
	})
	return events, err
}
