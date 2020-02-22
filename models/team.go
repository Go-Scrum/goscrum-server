package models

type Team struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
