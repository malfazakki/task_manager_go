package main

import (
	"log"
	"task-manager/config"
	"task-manager/internal/handler"
	"task-manager/internal/middleware"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"task-manager/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	e := echo.New()

	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	userRepo := repository.NewUserRepository(config.DB)

	taskRepo := repository.NewTaskRepository(config.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	taskRoutes := e.Group("/tasks", middleware.AuthMiddleware)
	taskRoutes.POST("", taskHandler.CreateTask)
	taskRoutes.GET("", taskHandler.GetTask)

	e.Logger.Fatal(e.Start(":8080"))
}
