package models

// Standuper model used for serialization/deserialization stored ChannelMembers
type Participant struct {
	Model
	WorkspaceID string     `db:"workspace_id" json:"workspace_id"`
	UserID      string     `db:"user_id" json:"user_id"`
	ChannelID   string     `db:"channel_id" json:"channel_id"`
	Role        string     `db:"role" json:"role"`
	RealName    string     `db:"real_name" json:"real_name"`
	ChannelName string     `db:"channel_name" json:"channel_name"`
	Projects    []*Project `gorm:"many2many:project_participants;"`
}
