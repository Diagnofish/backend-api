package service

import (
	"diagnofish/model"
	repo "diagnofish/repository"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type FishService interface {
	Detection(userFile *model.UserFile) (*model.FishDetection, error)
	StoreImage(userFile *model.UserFile) error
}

type fishService struct {
	fishRepository repo.FishRepository
}

func NewFishService(fishRepository repo.FishRepository) FishService {
	return &fishService{fishRepository}
}

func (f *fishService) Detection(userFile *model.UserFile) (*model.FishDetection, error) {
	var fishDetection model.FishDetection

	apiURL := "http://localhost:8000/detection"
	client := resty.New()

	// kirim file ke ML service
	resp, err := client.R().SetFile("image", userFile.FileDirectory).Post(apiURL)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), &fishDetection)
	if err != nil {
		return nil, err
	}
	return &fishDetection, nil
}

func (f *fishService) StoreImage(userFile *model.UserFile) error {
	// bucketName := "testing-capstone-environment"

	// ctx := context.Background()

	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	return err
	// }

	// bucket := client.Bucket(bucketName)
	// object := bucket.Object(userFile.Filename)
	// writer := object.NewWriter(ctx)
	// defer writer.Close()

	// file, err := os.Open(userFile.FileDirectory)
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

	if err := os.Remove(userFile.FileDirectory); err != nil {
		return err
	}

	return nil
}
