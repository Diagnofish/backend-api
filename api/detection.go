package api

import (
	"diagnofish/model"
	"diagnofish/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DetectionAPI interface {
	Detection(c *gin.Context)
}

type detectionAPI struct {
	detectionService service.DetectionService
}

func NewDetectionAPI(detectionService service.DetectionService) *detectionAPI {
	return &detectionAPI{detectionService}
}

func (f *detectionAPI) Detection(c *gin.Context) {
	var fishDetection model.FishDetection

	uuid := uuid.New()
	id := uuid.String()[:8]
	email, _ := c.Get("email")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("no such file"))
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("failed to get directory"))
	}

	filename := fmt.Sprintf("%s%s", id, filepath.Ext(file.Filename))
	fileDirectory := filepath.Join(dir, "images", filename)

	if err = c.SaveUploadedFile(file, fileDirectory); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("file failed to upload"))
		return
	}

	var imageData = model.ImageData{
		ID:            id,
		FileOwner:     fmt.Sprintf("%v", email),
		Filename:      filename,
		FileDirectory: fileDirectory,
	}

	fishDetection, err = f.detectionService.Detection(&imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	if err = f.detectionService.StoreImage(&imageData, &fishDetection); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, fishDetection)
}
