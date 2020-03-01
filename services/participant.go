package services

import (
	"goscrum/server/models"

	"github.com/jinzhu/gorm"
)

type ParticipantService struct {
	db *gorm.DB
}

func NewParticipantService(db *gorm.DB) ParticipantService {
	return ParticipantService{db: db}
}

func (service *ParticipantService) GetParticipantByUserId(userId string) (*models.Participant, error) {
	participant := models.Participant{}
	err := service.db.
		Preload("Projects").
		Where("user_id = ?", userId).
		First(&participant).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &participant, nil
}
