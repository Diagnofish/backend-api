package service

import (
	"diagnofish/model"
	repo "diagnofish/repository"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type DetectionService interface {
	Detection(imageData *model.ImageData) (model.DetectedFish, error)
	StoreImage(imageData *model.ImageData, detectedFish *model.DetectedFish) error
	GetList(userId string) ([]model.DetectedFish, error)
	GetByID(id string, userId string) (*model.DetectionDetail, error)
}

type detectionService struct {
	detectionRepository repo.DetectionRepository
}

func NewDetectionService(detectionRepository repo.DetectionRepository) DetectionService {
	return &detectionService{detectionRepository}
}

func (d *detectionService) Detection(imageData *model.ImageData) (model.DetectedFish, error) {
	var detectedFish model.DetectedFish

	apiURL := "https://prediction-api-hnobhrzdiq-et.a.run.app/detection"
	client := resty.New()

	// kirim file ke ML service
	resp, err := client.R().SetFile("image", imageData.FileDirectory).Post(apiURL)
	if err != nil {
		os.Remove(imageData.FileDirectory)
		return detectedFish, err
	}

	err = json.Unmarshal(resp.Body(), &detectedFish)
	if err != nil {
		os.Remove(imageData.FileDirectory)
		return detectedFish, err
	}

	detectedFish.ID = imageData.ID
	detectedFish.ImageFilename = imageData.Filename
	detectedFish.UserId = imageData.FileOwner

	return detectedFish, nil
}

func (d *detectionService) StoreImage(imageData *model.ImageData, detectedFish *model.DetectedFish) error {
	// bucketName := "diagnofish-bucket"

	// ctx := context.Background()

	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	os.Remove(imageData.FileDirectory)
	// 	return err
	// }

	// bucket := client.Bucket(bucketName)
	// object := bucket.Object("Fish Detection Image/" + imageData.Filename)
	// writer := object.NewWriter(ctx)
	// defer writer.Close()

	// file, err := os.Open(imageData.FileDirectory)
	// if err != nil {
	// 	// fmt.Printf("Failed to open file: %v", err)
	// 	os.Remove(imageData.FileDirectory)
	// 	return err
	// }
	// defer file.Close()

	// _, err = io.Copy(writer, file)
	// if err != nil {
	// 	// fmt.Printf("Failed to copy file to GCS: %v", err)
	// 	os.Remove(imageData.FileDirectory)
	// 	return err
	// }

	if err := os.Remove(imageData.FileDirectory); err != nil {
		return err
	}

	err := d.detectionRepository.Store(detectedFish)
	if err != nil {
		return err
	}

	return nil
}

func (d *detectionService) GetList(userId string) ([]model.DetectedFish, error) {
	history, err := d.detectionRepository.GetList(userId)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (d *detectionService) GetByID(id string, userId string) (*model.DetectionDetail, error) {
	detectionDetail, err := d.detectionRepository.GetByID(id, userId)
	if err != nil {
		return nil, err
	}

	return detectionDetail, nil
}
