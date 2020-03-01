package services

import (
	"goscrum/server/models"
	"time"

	"github.com/jinzhu/gorm"
)

type UserActivityService struct {
	db *gorm.DB
}

func NewUserActivityService(db *gorm.DB) UserActivityService {
	return UserActivityService{db: db}
}

func (service *UserActivityService) GetLastUserActivity(participantId string) (*models.UserActivity, error) {
	today := time.Now().Add(time.Second * time.Duration(5)) // TODO this is db issue
	yesterday := today.AddDate(0, 0, -1)
	userActivity := &models.UserActivity{}
	err := service.db.
		Where("participant_id = ? AND created_at BETWEEN ? AND ?", participantId, yesterday, today).
		Order("created_at desc").
		First(&userActivity).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return userActivity, err
}

func (service *UserActivityService) GetUserActivities(participantId string) ([]models.UserActivity, error) {
	today := time.Now().Add(time.Second * time.Duration(5)) // TODO this is db issue
	yesterday := today.AddDate(0, 0, -1)
	var userActivities []models.UserActivity
	err := service.db.
		Where("participant_id = ? AND created_at BETWEEN ? AND ?", participantId, yesterday, today).
		Order("created_at desc").
		Find(&userActivities).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return userActivities, err
}

func (service *UserActivityService) Add(userActivity models.UserActivity) error {
	return service.db.Create(&userActivity).Error
}

func (service *UserActivityService) DeleteActivitiesForUser(participantId string) error {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	return service.db.
		Where("participant_id = ? AND created_at BETWEEN ? AND ?", participantId, today, yesterday).
		Delete(models.UserActivity{}).Error
}
