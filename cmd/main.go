package main

import (
	"main/internal/cache"
	auth_controller "main/internal/controllers/http/v1/auth"
	"main/internal/pkg/config"
	"main/internal/pkg/postgres"
	auth "main/internal/repository/postgres/auth"
	auth_service "main/internal/services/auth"
	auth_use_case "main/internal/usecase/auth"
	send_sms_use_case "main/internal/usecase/send_sms"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// serverPost := ":" + config.GetConfig().Port
	// r.Run("0.0.0.0:" + config.GetConfig().Port)

	r := gin.Default()

	//databases
	postgresDB := postgres.NewDB()

	r.Static("/media", "./media")

	//cache
	newCache := cache.NewCache(config.GetConfig().RedisHost, config.GetConfig().RedisDB, time.Duration(config.GetConfig().RedisExpires)*time.Second)

	//repositories
	authRepository := auth.NewRepository(postgresDB)

	//usecase
	authUseCase := auth_use_case.NewUseCase(authRepository)
	sendSMSUseCase := send_sms_use_case.NewUseCase()
	// emailUseCase := email_use_case.NewUseCase()

	//services
	authService := auth_service.NewService(authRepository, newCache, sendSMSUseCase, authUseCase)

	//controller
	authController := auth_controller.NewController(authService)

	//middleware
	// authMiddleware := auth_middleware.NewMiddleware(authService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		v1 := api.Group("v1")

		// #auth
		// send otp
		v1.POST("/cabinet/login/send/otp", authController.SendOtp)

	}

	r.Run("0.0.0.0:" + config.GetConfig().Port)
}
