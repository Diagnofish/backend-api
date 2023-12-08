package service

import (
	"diagnofish/model"
	repo "diagnofish/repository"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type DetectionService interface {
	Detection(imageData *model.ImageData) (model.FishDetection, error)
	StoreImage(imageData *model.ImageData, fishDetection *model.FishDetection) error
}

type detectionService struct {
	detectionRepository repo.DetectionRepository
}

func NewDetectionService(detectionRepository repo.DetectionRepository) DetectionService {
	return &detectionService{detectionRepository}
}

func (d *detectionService) Detection(imageData *model.ImageData) (model.FishDetection, error) {
	var fishDetection model.FishDetection

	apiURL := "http://localhost:8000/detection"
	client := resty.New()

	// kirim file ke ML service
	resp, err := client.R().SetFile("image", imageData.FileDirectory).Post(apiURL)
	if err != nil {
		return fishDetection, err
	}

	err = json.Unmarshal(resp.Body(), &fishDetection)
	if err != nil {
		return fishDetection, err
	}

	fishDetection.ID = imageData.ID
	fishDetection.ImageFilename = imageData.Filename
	fishDetection.Email = imageData.FileOwner

	return fishDetection, nil
}

func (d *detectionService) StoreImage(imageData *model.ImageData, fishDetection *model.FishDetection) error {
	// bucketName := "testing-capstone-environment"

	// ctx := context.Background()

	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	return err
	// }

	// bucket := client.Bucket(bucketName)
	// object := bucket.Object(imageData.Filename)
	// writer := object.NewWriter(ctx)
	// defer writer.Close()

	// file, err := os.Open(imageData.FileDirectory)
	// if err != nil {
	// 	// fmt.Printf("Failed to open file: %v", err)
	// 	return err
	// }
	// defer file.Close()

	// _, err = io.Copy(writer, file)
	// if err != nil {
	// 	// fmt.Printf("Failed to copy file to GCS: %v", err)
	// 	return err
	// }

	if err := os.Remove(imageData.FileDirectory); err != nil {
		return err
	}

	err := d.detectionRepository.Store(fishDetection)
	if err != nil {
		return err
	}

	return nil
}