package service

type DetectionService interface {
	Detection()
}

func NewDetectionService() DetectionService {
	return nil
}
