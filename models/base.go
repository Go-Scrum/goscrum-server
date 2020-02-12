package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Model struct {
	ID        string     `gorm:"primary_key" json:"id" valid:"uuidv4, optional"`
	CreatedAt time.Time  `valid:"-" json:"created_at"`
	UpdatedAt time.Time  `valid:"-" json:"-"`
	DeletedAt *time.Time `sql:"index" valid:"-" json:"-"`
}

//BeforeUpdate is a hook to set the created_at column to UNIX timestamp int.
func (m *Model) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedAt", time.Now())
}

//BeforeCreate is a hook to set the created_at column to UNIX timestamp int.
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	u4 := uuid.NewV4()
	err := scope.SetColumn("ID", u4.String())

	if err != nil {
		return err
	}

	err = scope.SetColumn("UpdatedAt", time.Now())

	if err != nil {
		return err
	}

	return scope.SetColumn("CreatedAt", time.Now())
}
