package repository

import (
	"diagnofish/model"

	"gorm.io/gorm"
)

type FishRepository interface {
	Store(Fish *model.FishDetection) error
}

type fishRepository struct {
	db *gorm.DB
}

func NewFishRepo(db *gorm.DB) *fishRepository {
	return &fishRepository{db}
}

func (f *fishRepository) Store(Fish *model.FishDetection) error {
	if err := f.db.Create(Fish).Error; err != nil {
		return err
	}

	return nil
}
