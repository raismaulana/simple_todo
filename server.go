package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/simple_todo/config"
	"github.com/raismaulana/simple_todo/controller"
	"github.com/raismaulana/simple_todo/middleware"
	"github.com/raismaulana/simple_todo/repository"
	"github.com/raismaulana/simple_todo/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.StaticUserRepository(db)
	jwtService     service.JWTService        = service.JWTAuthService()
	authService    service.AuthService       = service.StaticAuthService(userRepository)
	userService    service.UserService       = service.StaticUserService(userRepository)
	authController controller.AuthController = controller.StaticAuthController(authService, jwtService)
	userController controller.UserController = controller.StaticUserController(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	e, err := casbin.NewEnforcer("config/authz-model.conf", "config/authz-policy.csv")
	if err != nil {
		panic(err)
	}
	log.Println(e)
	r := gin.Default()
	guestRoutes := r.Group("api/auth")
	{
		guestRoutes.POST("/login", authController.Login)
		guestRoutes.POST("/register", authController.Register)
	}
	authRoutes := r.Group("api", middleware.AuthorizeJWT(jwtService))
	authRoutes.Use(middleware.AuthorizerRole(e, jwtService, authService))
	{
		authRoutes.GET("/user/get", userController.GetAllUser)
		authRoutes.PUT("/user/update", userController.UpdateUser)
	}

	r.Run()
}
