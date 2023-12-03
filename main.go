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
	UserAPIHandler api.UserAPI
	FishAPIHandler api.FishAPI
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
	sessionRepo := repo.NewSessionRepo(db)
	fishRepo := repo.NewFishRepo(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	fishService := service.NewFishService(fishRepo)

	userAPIHandler := api.NewUserAPI(userService)
	fishAPIHandler := api.NewFishAPI(fishService)

	apiHandler := APIHandler{
		UserAPIHandler: userAPIHandler,
		FishAPIHandler: fishAPIHandler,
	}

	user := gin.Group("/user")
	{
		user.POST("/register", apiHandler.UserAPIHandler.Register)
		user.POST("/login", apiHandler.UserAPIHandler.Login)
	}

	// detection := gin.Group("/fish", middleware.Auth())
	detection := gin.Group("/fish")
	{
		detection.POST("/detection", apiHandler.FishAPIHandler.Detection)
	}

	return gin
}
