package main

import (
	cache "main/internal/cache"
	auth_controller "main/internal/controllers/http/v1/auth"
	order_controller "main/internal/controllers/http/v1/order"
	price_controller "main/internal/controllers/http/v1/price"
	user_controller "main/internal/controllers/http/v1/user"
	auth_middleware "main/internal/middleware/auth"
	"main/internal/pkg/config"
	"main/internal/pkg/postgres"
	"main/internal/repository/postgres/order"
	"main/internal/repository/postgres/price"
	"main/internal/repository/postgres/user"
	"main/internal/services/auth"
	"main/internal/services/email"
	file_service "main/internal/services/file"
	order_service "main/internal/services/order"
	price_service "main/internal/services/price"
	user_service "main/internal/services/user"
	"main/internal/services/ws"
	auth_use_case "main/internal/usecase/auth"
	order_use_case "main/internal/usecase/order"
	price_use_case "main/internal/usecase/price"
	user_use_case "main/internal/usecase/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	serverPost := ":" + config.GetConfig().Port

	r := gin.Default()

	//databases
	postgresDB := postgres.NewDB()

	r.Static("/media", "./media")

	//cache
	newCache := cache.NewCache(config.GetConfig().RedisHost, config.GetConfig().RedisDB, time.Duration(config.GetConfig().RedisExpires)*time.Second)

	//repositories
	userRepository := user.NewRepository(postgresDB)
	orderRepository := order.NewRepository(postgresDB)
	priceRepository := price.NewRepository(postgresDB)

	// ws
	wsManager := ws.NewManager(orderRepository)
	r.GET("/ws", wsManager.HandleDriverWS)
	r.GET("/ws/client", wsManager.HandleClientWS)

	//services
	// videoService := video_service.NewService()
	// audioService := audio_service.NewService()
	authService := auth.NewService(userRepository)
	emailService := email.NewEmailService()
	fileService := file_service.NewService()
	userService := user_service.NewService(userRepository)
	orderService := order_service.NewService(orderRepository)
	priceService := price_service.NewService(priceRepository)

	//usecase
	authUseCase := auth_use_case.NewUseCase(authService, newCache, emailService, userService)
	userUseCase := user_use_case.NewUseCase(userService, authService, fileService)
	orderUseCase := order_use_case.NewUseCase(orderService, authService, wsManager, newCache)
	priceUseCase := price_use_case.NewUseCase(priceService)

	//controller
	authController := auth_controller.NewController(authUseCase)
	userController := user_controller.NewController(userUseCase)
	orderController := order_controller.NewController(orderUseCase)
	priceController := price_controller.NewController(priceUseCase)

	//middleware
	authMiddleware := auth_middleware.NewMiddleware(authService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://192.168.1.120:8081", "http://172.20.10.5:8081"},
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
		//forgot password
		v1.POST("/send/code", authController.SendEmailCode)
		//check code
		v1.POST("/check/code", authController.CheckCode)
		//resend code
		v1.POST("/resend/code", authController.ResendCode)

		// #user
		// update
		v1.PATCH("/user/update", authMiddleware.AuthMiddleware(), userController.Update)
		// delete
		v1.DELETE("/user/delete", authMiddleware.AuthMiddleware(), userController.Delete)

		// #order
		// create
		v1.POST("/order/create", authMiddleware.AuthMiddleware(), orderController.CreateOrder)
		// get list
		v1.GET("/order/list", authMiddleware.AuthMiddleware(), orderController.GetList)

		// #driver
		v1.PATCH("/accept/order/:uuid", authMiddleware.AuthMiddleware(), orderController.OrderAccept)

		// #price
		v1.GET("/get/price", authMiddleware.AuthMiddleware(), priceController.GetPriceByLocation)

	}

	r.Run(serverPost)

}
