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
	db *gorm.DB = config.ConnectDB()

	//Repository..........
	userRepository repository.UserRepository = repository.NewUserConnection(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)

	//Service......
	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)
	bookService service.BookService = service.NewBookService(bookRepository)
	authService service.AuthService = service.NewAuthService(userRepository)

	//Controller........
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDB(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)       // login
		authRoutes.POST("/register", authController.Register) // register
	}

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/profile", userController.Profile) // Get User
		userRoutes.PUT("/profile", userController.Update)  // Update User
	}

	bookRoutes := r.Group("api/book")
	{
		bookRoutes.GET("/", bookController.GetAll)       // Get All Book
		bookRoutes.POST("/", bookController.Insert)      // Create a Book
		bookRoutes.GET("/:id", bookController.FindById)  // Get Book by ID
		bookRoutes.PUT("/:id", bookController.Update)    // Update a Book based on Id
		bookRoutes.DELETE("/:id", bookController.Delete) // Delete Book

	}
	r.Run()
}
