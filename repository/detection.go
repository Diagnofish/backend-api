package repository

import (
	"diagnofish/model"

	"gorm.io/gorm"
)

type DetectionRepository interface {
	Store(DetectionResult *model.DetectedFish) error
	GetList(userId string) ([]model.DetectedFish, error)
	GetByID(id string, userId string) (*model.DetectedFish, error)
}

type detectionRepository struct {
	db *gorm.DB
}

func NewFishRepo(db *gorm.DB) *detectionRepository {
	return &detectionRepository{db}
}

func (d *detectionRepository) Store(DetectionResult *model.DetectedFish) error {
	if err := d.db.Create(DetectionResult).Error; err != nil {
		return err
	}

	return nil
}

func (d *detectionRepository) GetList(userId string) ([]model.DetectedFish, error) {
	var detectedFish model.DetectedFish
	var history = []model.DetectedFish{}

	rows, err := d.db.Model(&detectedFish).Select("*").Where("user_id = ?", userId).Rows()
	if err != nil {
		return history, err
	}
	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &history)
	}

	return history, nil
}

func (d *detectionRepository) GetByID(id string, userId string) (*model.DetectedFish, error) {
	var detectionData model.DetectedFish

	if err := d.db.Where("id = ? AND user_id = ?", id, userId).First(&detectionData).Error; err != nil {
		return nil, err
	}

	return &detectionData, nil
}
