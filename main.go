package main

import (
	"diagnofish/api"
	"diagnofish/db"
	"diagnofish/middleware"
	"diagnofish/model"
	repo "diagnofish/repository"
	"diagnofish/service"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler      api.UserAPI
	DetectionAPIHandler api.DetectionAPI
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	db := db.NewDB()

	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("[%s] \"%s %s %s\"\n",
	// 		param.TimeStamp.Format(time.RFC822),
	// 		param.Method,
	// 		param.Path,
	// 		param.ErrorMessage,
	// 	)
	// }))
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

	conn.AutoMigrate(&model.User{}, &model.DetectedFish{}, &model.Session{})

	router = RunServer(conn, router)

	fmt.Println("Server is running on port 8080")
	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	userRepo := repo.NewUserRepo(db)
	sessionRepo := repo.NewSessionRepo(db)
	detectionRepo := repo.NewFishRepo(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	detectionService := service.NewDetectionService(detectionRepo)

	userAPIHandler := api.NewUserAPI(userService)
	detectionAPIHandler := api.NewDetectionAPI(detectionService)

	apiHandler := APIHandler{
		UserAPIHandler:      userAPIHandler,
		DetectionAPIHandler: detectionAPIHandler,
	}

	user := gin.Group("/user")
	{
		user.POST("/register", apiHandler.UserAPIHandler.Register)
		user.POST("/login", apiHandler.UserAPIHandler.Login)
		user.POST("/logout", apiHandler.UserAPIHandler.Logout)
	}

	detection := gin.Group("/detection", middleware.Auth())
	// detection := gin.Group("/fish")
	{
		detection.POST("", apiHandler.DetectionAPIHandler.Detection)
		detection.GET("/history", apiHandler.DetectionAPIHandler.GetList)
		detection.GET("/history/:id", apiHandler.DetectionAPIHandler.GetByID)
	}

	return gin
}
