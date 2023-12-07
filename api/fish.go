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
		Email:         fmt.Sprintf("%v", email),
		Filename:      filename,
		FileDirectory: fileDirectory,
	}

	fishDetection, err = f.fishService.Detection(&imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	if err = f.fishService.StoreImage(&imageData, &fishDetection); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	// fmt.Println(resp.String())
	// fmt.Println("respp")
	// fmt.Println(respp)

	// c.JSON(resp.StatusCode(), respp)

	c.JSON(http.StatusOK, fishDetection)

	// forwardImage(c, filename, fileDirectory)
	// uploadToGCS(filename, fileDirectory)

	// upload image

	// meneruskan ke ML
	// upload ke bucket
	// mendapat response dari ML

	// insert ke database
	// tulis response ke client
	// delete local image
}
