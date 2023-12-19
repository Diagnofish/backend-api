package main

import (
	"diagnofish/api"
	"diagnofish/db"
	"diagnofish/middleware"
	"diagnofish/model"
	repo "diagnofish/repository"
	"diagnofish/service"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler      api.UserAPI
	DetectionAPIHandler api.DetectionAPI
}

func main() {
	gin.SetMode(gin.DebugMode)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.Default()
		db := db.NewDB()

		router.Use(gin.Recovery())

		dbCredential := model.Credential{
			Host:         "34.128.122.77",
			Username:     "postgres",
			Password:     "db_diagnofish",
			DatabaseName: "diagnofish_db",
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

	}()

	wg.Wait()
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
