package main

import (
	"diagnofish/api"
	"diagnofish/db"
	"diagnofish/model"
	repo "diagnofish/repository"
	"diagnofish/service"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler      api.UserAPI
	DetectionAPIHandler api.DetectionAPI
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	db := db.NewDB()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] \"%s %s %s\"\n",
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "123",
		DatabaseName: "diagnofish",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{})

	router = RunServer(conn, router)

	fmt.Println("Server is running on port 8080")
	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	userRepo := repo.NewUserRepo(db)

	userService := service.NewUserService(userRepo)
	detectionService := service.NewDetectionService

	userAPIHandler := api.NewUserAPI(userService)
	detectionAPIHandler := api.NewDetectionAPI(detectionService())

	apiHandler := APIHandler{
		UserAPIHandler:      userAPIHandler,
		DetectionAPIHandler: detectionAPIHandler,
	}

	user := gin.Group("/user")
	{
		user.POST("/register", apiHandler.UserAPIHandler.Register)
		// user.POST("/login", apiHandler.UserAPIHandler.Login)
	}

	detection := gin.Group("/detection")
	{
		detection.POST("/predict", apiHandler.DetectionAPIHandler.Predict)
	}

	return gin
}
