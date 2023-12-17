package repository

import (
	"diagnofish/model"

	"gorm.io/gorm"
)

type DetectionRepository interface {
	Store(DetectionResult *model.DetectedFish) error
	GetList(userId string) ([]model.DetectedFish, error)
	GetByID(id string, userId string) (*model.DetectionDetail, error)
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

func (d *detectionRepository) GetByID(id string, userId string) (*model.DetectionDetail, error) {
	var detectionDetail model.DetectionDetail

	err := d.db.Table("detected_fishes").
		Select(
			"detected_fishes.id, detected_fishes.image_filename, detected_fishes.fish_name, detected_fishes.result, detected_fishes.confidence_score, class_details.description, class_details.symptom, class_details.cause, class_details.treatment, class_details.prevention").
		Joins(
			"JOIN class_details ON detected_fishes.result = class_details.result").
		Where("detected_fishes.id = ? AND detected_fishes.user_id = ?", id, userId).
		First(&detectionDetail).Error
	if err != nil {
		return nil, err
	}

	return &detectionDetail, nil

	// var detectionData model.DetectedFish

	// if err := d.db.Where("id = ? AND user_id = ?", id, userId).First(&detectionData).Error; err != nil {
	// 	return nil, err
	// }

	// return &detectionData, nil
}
