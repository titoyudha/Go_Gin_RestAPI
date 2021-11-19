package main

import (
	"github.com/gin-gonic/gin"
	"github.com/titoyudha/Go_Gin_RestAPI/config"
	"github.com/titoyudha/Go_Gin_RestAPI/controller"
	"github.com/titoyudha/Go_Gin_RestAPI/repository"
	"github.com/titoyudha/Go_Gin_RestAPI/service"

	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectDB()
	userRepository repository.UserRepository = repository.NewUserConnection(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDB(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("register", authController.Register)
	}

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	r.Run()
}
