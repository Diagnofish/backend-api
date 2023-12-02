package api

import (
	"diagnofish/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DetectionAPI interface {
	Predict(c *gin.Context)
}

type detectionAPI struct {
	detectionService service.DetectionService
}

func NewDetectionAPI(detectionService service.DetectionService) *detectionAPI {
	return &detectionAPI{detectionService}
}

func (p *detectionAPI) Predict(c *gin.Context) {
	u := uuid.New()

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fdsaf"})
		return
	}

	filename := fmt.Sprintf("%s%s", u.String()[:8], filepath.Ext(file.Filename))

	fileLocation := filepath.Join(dir, "images", filename)

	err = c.SaveUploadedFile(file, fileLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
