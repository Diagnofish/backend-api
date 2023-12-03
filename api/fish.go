package api

import (
	"diagnofish/model"
	"diagnofish/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type FishAPI interface {
	Detection(c *gin.Context)
}

type fishAPI struct {
	fishService service.FishService
}

func NewFishAPI(fishService service.FishService) *fishAPI {
	return &fishAPI{fishService}
}

func (f *fishAPI) Detection(c *gin.Context) {
	//result, err := f.fishService.Detection()

	uploadingImage(c)

	// upload image

	// meneruskan ke ML
	// mendapat response dari ML

	// insert ke database
	// tulis response ke client
	// delete local image
}

func uploadingImage(c *gin.Context) {
	uuid := uuid.New()

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("failed to get directory"))
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("no such file"))
		return
	}

	filename := fmt.Sprintf("%s%s", uuid.String()[:8], filepath.Ext(file.Filename))

	fileDir := filepath.Join(dir, "images", filename)

	err = c.SaveUploadedFile(file, fileDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("file failed to upload"))
		return
	}

	forwardImage(c, filename, fileDir)

	// c.JSON(http.StatusOK, model.NewSuccessResponse("file uploaded successfully"))
}

func forwardImage(c *gin.Context, filename string, path string) {
	apiURL := "http://localhost:8000/detection"

	client := resty.New()

	resp, err := client.R().SetFile("image", path).Post(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(resp.StatusCode(), gin.H{
		"path":     path,
		"response": resp.String(),
	})

	// c.JSON(http.StatusInternalServerError, model.NewErrorResponse("invalid decode json from other service"))

	// if err := c.BindJSON(resp); err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.NewErrorResponse("invalid decode json from other service"))
	// 	return
	// }
}
