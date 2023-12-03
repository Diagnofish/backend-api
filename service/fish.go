package service

import (
	"diagnofish/model"
	repo "diagnofish/repository"
)

type FishService interface {
	Detection() (*model.FishDetection, error)
}

type fishService struct {
	fishRepository repo.FishRepository
}

func NewFishService(fishRepository repo.FishRepository) FishService {
	return &fishService{fishRepository}
}

func (f *fishService) Detection() (*model.FishDetection, error) {

	// apiURL := ""

	// client := resty.New()

	// resp, err := client.R().
	//     SetFile("profile_img", "/Users/jeeva/test-img.png").
	//     Post(apiURL)
	// if err != nil {
	// 	return nil, model.NewErrorResponse("email has already exists")
	// }

	// fishDetection := model.FishDetection{

	// }

	return nil, nil
}
