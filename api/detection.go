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
	GetList(c *gin.Context)
	GetByID(c *gin.Context)
}

type detectionAPI struct {
	detectionService service.DetectionService
}

func NewDetectionAPI(detectionService service.DetectionService) *detectionAPI {
	return &detectionAPI{detectionService}
}

func (d *detectionAPI) Detection(c *gin.Context) {
	var detectedFish model.DetectedFish

	uuid := uuid.New()
	id := uuid.String()[:8]
	userId, _ := c.Get("user_id")

	fmt.Println(fmt.Sprintf("%v", userId))

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("no such file"))
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("failed to get directory"))
	}

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(file.Filename)

	validExtension := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("only image files (jpg, jpeg, png) are allowed"))
		return
	}

	filename := fmt.Sprintf("%s%s", id, filepath.Ext(file.Filename))
	fileDirectory := filepath.Join(dir, "images", filename)

	if err = c.SaveUploadedFile(file, fileDirectory); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("file failed to upload"))
		return
	}

	var imageData = model.ImageData{
		ID:            id,
		FileOwner:     fmt.Sprintf("%v", userId),
		Filename:      filename,
		FileDirectory: fileDirectory,
	}

	detectedFish, err = d.detectionService.Detection(&imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	if err = d.detectionService.StoreImage(&imageData, &detectedFish); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, detectedFish)
}

func (d *detectionAPI) GetList(c *gin.Context) {
	uuid, _ := c.Get("user_id")
	userId := fmt.Sprintf("%v", uuid)

	history, err := d.detectionService.GetList(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, history)
}

func (d *detectionAPI) GetByID(c *gin.Context) {
	detectionID := c.Param("id")
	uuid, _ := c.Get("user_id")
	userId := fmt.Sprintf("%v", uuid)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid task id"))
	// 	return
	// }

	dataDetection, err := d.detectionService.GetByID(detectionID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	c.JSON(http.StatusOK, dataDetection)
}
