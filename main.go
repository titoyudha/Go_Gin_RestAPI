package main

import (
	"github.com/gin-gonic/gin"
	"github.com/titoyudha/Go_Gin_RestAPI/config"
	"github.com/titoyudha/Go_Gin_RestAPI/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectDB()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
	}
	r.Run()
}
