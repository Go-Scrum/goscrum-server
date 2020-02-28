package models

import "time"

type RequestCommits struct {
	Since time.Time `json:"since"`
}

type RequestIssues struct {
	State            string    `json:"state"`
	AuthorID         int       `json:"author_id"`
	AuthorUsername   string    `json:"author_username"`
	AssigneeID       int       `json:"assignee_id"`
	AssigneeUsername string    `json:"assignee_username"`
	UpdatedAfter     time.Time `json:"updated_after"`
	Search           string    `json:"search"`
}
