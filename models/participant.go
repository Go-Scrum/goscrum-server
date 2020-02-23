package models

// Standuper model used for serialization/deserialization stored ChannelMembers
type Participant struct {
	Model
	WorkspaceID string     `db:"workspace_id" json:"workspace_id"`
	UserID      string     `db:"user_id" json:"user_id"`
	Role        string     `db:"role" json:"role"`
	RealName    string     `db:"real_name" json:"real_name"`
	FirstName   string     `db:"first_name" json:"first_name"`
	LastName    string     `db:"last_name" json:"last_name"`
	NickName    string     `db:"nick_name" json:"nick_name"`
	Email       string     `db:"email" json:"channel_name"`
	Projects    []*Project `gorm:"many2many:project_participants;"`
}
