package repository

import (
	"diagnofish/model"

	"gorm.io/gorm"
)

type DetectionRepository interface {
	Store(DetectionResult *model.FishDetection) error
}

type detectionRepository struct {
	db *gorm.DB
}

func NewFishRepo(db *gorm.DB) *detectionRepository {
	return &detectionRepository{db}
}

func (d *detectionRepository) Store(DetectionResult *model.FishDetection) error {
	if err := d.db.Create(DetectionResult).Error; err != nil {
		return err
	}

	return nil
}
